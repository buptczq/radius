package radius

import (
	"context"
	"errors"
	"net"
	"sync"
)

type packetResponseWriter struct {
	// listener that received the packet
	conn net.PacketConn
	addr net.Addr
}

func (r *packetResponseWriter) Write(packet *Packet) error {
	raw, err := packet.Encode()
	if err != nil {
		return err
	}
	if _, err := r.conn.WriteTo(raw, r.addr); err != nil {
		return err
	}
	return nil
}

// PacketServer listens for RADIUS requests on a packet-based protocols (e.g.
// UDP).
type PacketServer struct {
	// The address on which the server listens. Defaults to :1812.
	Addr string
	// The network on which the server listens. Defaults to udp.
	Network      string
	SecretSource SecretSource
	Handler      Handler

	// Skip incoming packet authenticity validation.
	// This should only be set to true for debugging purposes.
	InsecureSkipVerify bool

	mu          sync.Mutex
	ctx         context.Context
	ctxDone     context.CancelFunc
	wg          sync.WaitGroup
	connections map[net.PacketConn]struct{}
	doneChan    chan struct{}
}

func (s *PacketServer) getDoneChanLocked() chan struct{} {
	if s.doneChan == nil {
		s.doneChan = make(chan struct{})
	}
	return s.doneChan
}

func (s *PacketServer) getDoneChan() <-chan struct{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.getDoneChanLocked()
}

func (s *PacketServer) closeDoneChanLocked() {
	ch := s.getDoneChanLocked()
	select {
	case <-ch:
		// Already closed. Don't close again.
	default:
		// Safe to close here. We're the only closer, guarded
		// by s.mu.
		close(ch)
	}
}

func (s *PacketServer) closeConnectionsLocked() error {
	var err error
	for ln := range s.connections {
		if cerr := ln.Close(); cerr != nil && err == nil {
			err = cerr
		}
		delete(s.connections, ln)
	}
	return err
}

func (s *PacketServer) trackConnection(ln net.PacketConn, add bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.connections == nil {
		s.connections = make(map[net.PacketConn]struct{})
	}
	if add {
		// If the *Server is being reused after a previous
		// Close or Shutdown, reset its doneChan:
		if len(s.connections) == 0 {
			s.doneChan = nil
		}
		s.connections[ln] = struct{}{}
	} else {
		delete(s.connections, ln)
	}
}

// TODO: logger on PacketServer

// Serve accepts incoming connections on conn.
func (s *PacketServer) Serve(conn net.PacketConn) error {
	if s.Handler == nil {
		return errors.New("radius: nil Handler")
	}
	if s.SecretSource == nil {
		return errors.New("radius: nil SecretSource")
	}

	s.trackConnection(conn, true)
	defer s.trackConnection(conn, false)

	for {
		var buff [MaxPacketLength]byte
		n, remoteAddr, err := conn.ReadFrom(buff[:])
		if err != nil {
			select {
			case <-s.getDoneChan():
				return nil
			default:
			}
			if ne, ok := err.(net.Error); ok && !ne.Temporary() {
				return err
			}
			continue
		}

		buffCopy := make([]byte, n)
		copy(buffCopy, buff[:n])

		s.wg.Add(1)
		go func(buff []byte, remoteAddr net.Addr) {
			defer s.wg.Done()
			secret, err := s.SecretSource.RADIUSSecret(s.ctx, remoteAddr)
			if err != nil {
				// TODO: log only if server is not shutting down?
				return
			}
			if len(secret) == 0 {
				return
			}

			if !s.InsecureSkipVerify && !IsAuthenticRequest(buff, secret) {
				// TODO: log?
				return
			}

			packet, err := Parse(buff, secret)
			if err != nil {
				// TODO: error logger
				return
			}

			response := packetResponseWriter{
				conn: conn,
				addr: remoteAddr,
			}

			request := Request{
				LocalAddr:  conn.LocalAddr(),
				RemoteAddr: remoteAddr,
				Packet:     packet,
				ctx:        s.ctx,
			}

			s.Handler.ServeRADIUS(&response, &request)
		}(buffCopy, remoteAddr)
	}
}

// ListenAndServe starts a RADIUS server on the address given in s.
func (s *PacketServer) ListenAndServe() error {
	if s.Handler == nil {
		return errors.New("radius: nil Handler")
	}
	if s.SecretSource == nil {
		return errors.New("radius: nil SecretSource")
	}
	addrStr := ":1812"
	if s.Addr != "" {
		addrStr = s.Addr
	}
	network := "udp"
	if s.Network != "" {
		network = s.Network
	}
	s.ctx, s.ctxDone = context.WithCancel(context.Background())

	pc, err := net.ListenPacket(network, addrStr)
	if err != nil {
		return err
	}
	return s.Serve(pc)
}

// Shutdown gracefully stops the server. It first closes all connections (which
// stops accepting new packets) and then waits for running handlers to complete.
//
// Shutdown returns after all handlers have completed, or when ctx is canceled.
// The PacketServer is ready for re-use once the function returns nil.
func (s *PacketServer) Shutdown(ctx context.Context) error {
	s.mu.Lock()
	lnerr := s.closeConnectionsLocked()
	s.closeDoneChanLocked()
	s.mu.Unlock()

	if s.ctxDone != nil {
		s.ctxDone()
	}

	ch := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		return lnerr
	case <-ctx.Done():
		return ctx.Err()
	}
}
