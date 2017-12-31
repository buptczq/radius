package radius_test

import (
	"context"
	"fmt"
	"github.com/buptczq/radius"
	"testing"
	"time"
)

func TestPacketServer_basic(t *testing.T) {
	secret := []byte("123456790")

	server := radius.PacketServer{
		Addr:         "127.0.0.1:1812",
		Network:      "udp",
		SecretSource: radius.StaticSecretSource(secret),
		Handler: radius.HandlerFunc(func(w radius.ResponseWriter, r *radius.Request) {
			username := r.Packet.GetByName("User-Name").(string)
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
		packet.SetByName("User-Name", "tim")
		client := radius.Client{
			Retry: time.Millisecond * 50,
		}
		response, err := client.Exchange(context.Background(), packet, "127.0.0.1:1812")
		if err != nil {
			clientErr = err
			return
		}
		if response.Code != radius.CodeAccessAccept {
			clientErr = fmt.Errorf("expected CodeAccessAccept, got %s\n", response.Code)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		t.Fatal(err)
	}

	if clientErr != nil {
		t.Fail()
	}
}
