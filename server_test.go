package radius_test

import (
	"context"
	"fmt"
	"github.com/buptczq/go-radius"
	"net"
	"testing"
	"time"
)

func TestPacketServer_basic(t *testing.T) {
	addr, err := net.ResolveUDPAddr("udp", "localhost:1812")
	if err != nil {
		t.Fatal(err)
	}
	pc, err := net.ListenUDP("udp", addr)
	if err != nil {
		t.Fatal(err)
	}

	secret := []byte("123456790")

	server := radius.PacketServer{
		SecretSource: radius.StaticSecretSource(secret),
		Dictionary:   radius.Builtin,
		Handler: radius.HandlerFunc(func(w radius.ResponseWriter, r *radius.Request) {
			username := r.Packet.Value("User-Name").(string)
			if username == "tim" {
				w.Write(r.Response(radius.CodeAccessAccept))
			} else {
				w.Write(r.Response(radius.CodeAccessReject))
			}
		}),
	}

	var clientErr error
	go func() {
		defer server.Shutdown(context.Background())

		packet := radius.New(radius.CodeAccessRequest, secret)
		packet.Set("User-Name", "tim")
		client := radius.Client{
			Retry: time.Millisecond * 50,
		}
		response, err := client.Exchange(context.Background(), packet, pc.LocalAddr().String())
		if err != nil {
			clientErr = err
			return
		}
		if response.Code != radius.CodeAccessAccept {
			clientErr = fmt.Errorf("expected CodeAccessAccept, got %s\n", response.Code)
		}
	}()

	if err := server.Serve(pc); err != nil {
		t.Fatal(err)
	}

	server.Shutdown(context.Background())

	if clientErr != nil {
		t.Fatal(err)
	}
}
