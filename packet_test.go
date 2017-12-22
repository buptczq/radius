package radius_test

import (
	"bytes"
	"github.com/buptczq/radius"
	"net"
	"testing"
)

func Test_RFC2865_7_1(t *testing.T) {
	// Source: https://tools.ietf.org/html/rfc2865#section-7.1

	secret := []byte("xyzzy5461")

	// Request
	request := []byte{
		0x01, 0x00, 0x00, 0x38, 0x0f, 0x40, 0x3f, 0x94, 0x73, 0x97, 0x80, 0x57, 0xbd, 0x83, 0xd5, 0xcb,
		0x98, 0xf4, 0x22, 0x7a, 0x01, 0x06, 0x6e, 0x65, 0x6d, 0x6f, 0x02, 0x12, 0x0d, 0xbe, 0x70, 0x8d,
		0x93, 0xd4, 0x13, 0xce, 0x31, 0x96, 0xe4, 0x3f, 0x78, 0x2a, 0x0a, 0xee, 0x04, 0x06, 0xc0, 0xa8,
		0x01, 0x10, 0x05, 0x06, 0x00, 0x00, 0x00, 0x03,
	}

	p, err := radius.Parse(request, secret)
	if err != nil {
		t.Fatal(err)
	}
	if p.Code != radius.CodeAccessRequest {
		t.Fatal("expecting Code = PacketCodeAccessRequest")
	}
	if p.Identifier != 0 {
		t.Fatal("expecting Identifier = 0")
	}
	if p.Len() != 4 {
		t.Fatal("expecting 4 attributes")
	}
	if p.Value("User-Name").(string) != "nemo" {
		t.Fatal("expecting User-Name = nemo")
	}
	password, _ := radius.UserPassword(p.GetRaw(2), p.Secret, p.Authenticator[:])
	if string(password) != "arctangent" {
		t.Fatal("expecting User-Password = arctangent")
	}
	if !p.Value("NAS-IP-Address").(net.IP).Equal(net.ParseIP("192.168.1.16")) {
		t.Fatal("expecting NAS-IP-Address = 192.168.1.16")
	}
	if p.Value("NAS-Port").(uint32) != 3 {
		t.Fatal("expecting NAS-Port = 3")
	}

	{
		wire, err := p.Encode()
		if err != nil {
			t.Fatal(err)
		}
		if !RADIUSPacketsEqual(wire, request) {
			t.Fatal("expecting q.Encode() and request to be equal")
		}
	}

}

// RADIUSPacketsEqual returns if two RADIUS packets are equal, ignoring the
// order of attributes of different types.
func RADIUSPacketsEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if !bytes.Equal(a[:4], b[:4]) {
		return false
	}

	// hash is going to be different, as the attribute order could change

	aa, err := radius.ParseAttributes(a[20:])
	if err != nil {
		panic(err)
	}
	ab, err := radius.ParseAttributes(b[20:])
	if err != nil {
		panic(err)
	}

	if len(aa) != len(ab) {
		return false
	}

	for typeA, attrsA := range aa {
		if len(attrsA) != len(ab[typeA]) {
			return false
		}
		for i, attrA := range attrsA {
			if !bytes.Equal([]byte(attrA), []byte(ab[typeA][i])) {
				return false
			}
		}
	}

	return true
}

func Test_Vendor_Encode(t *testing.T) {
	// Aruba
	radius.Builtin.RegisterVendor("Aruba-Essid-Name", 14823, 5, radius.AttributeString)
	p := radius.New(radius.CodeAccessRequest, []byte("123"))
	p.Set("Aruba-Essid-Name", "wireless")
	wired, err := p.Encode()
	if err != nil {
		t.Fatal(err)
	}
	p, err = radius.Parse(wired, []byte("123"))
	if err != nil {
		t.Fatal(err)
	}
	if p.Code != radius.CodeAccessRequest {
		t.Fatal("expecting Code = PacketCodeAccessRequest")
	}
	if p.Value("Aruba-Essid-Name").(string) != "wireless" {
		t.Fatal("expecting Aruba-Essid-Name = wireless")
	}
}
