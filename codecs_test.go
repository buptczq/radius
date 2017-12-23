package radius

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func TestAttributeInteger(t *testing.T) {
	wired, err := Codecs[AttributeInteger].Encode(uint32(123456))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte{0, 1, 226, 64}) {
		t.Fail()
	}
	integer, err := Codecs[AttributeInteger].Decode([]byte{255, 255, 255, 255})
	if err != nil {
		t.Fatal(err)
	}
	if integer.(uint32) != 4294967295 {
		t.Fail()
	}
	ai := attributeInteger{}
	if v, _ := ai.Transform(uint(12)); v.(uint32) != 12 {
		t.Fail()
	}
	if v, _ := ai.Transform(int(12)); v.(uint32) != 12 {
		t.Fail()
	}
}

func TestAttributeInteger_Error(t *testing.T) {
	_, err := Codecs[AttributeInteger].Encode(uint64(123456))
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeInteger].Decode([]byte{255, 255, 255})
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeInteger64(t *testing.T) {
	wired, err := Codecs[AttributeInteger64].Encode(uint64(123456))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte{0, 0, 0, 0, 0, 1, 226, 64}) {
		t.Fail()
	}
	integer, err := Codecs[AttributeInteger64].Decode([]byte{255, 255, 255, 255, 255, 255, 255, 255})
	if err != nil {
		t.Fatal(err)
	}
	if integer.(uint64) != 18446744073709551615 {
		t.Fail()
	}
	ai := attributeInteger64{}
	if v, _ := ai.Transform(uint(12)); v.(uint64) != 12 {
		t.Fail()
	}
	if v, _ := ai.Transform(int(12)); v.(uint64) != 12 {
		t.Fail()
	}
}

func TestAttributeInteger64_Error(t *testing.T) {
	_, err := Codecs[AttributeInteger64].Encode(uint32(123456))
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeInteger64].Decode([]byte{255, 255, 255})
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeSigned(t *testing.T) {
	wired, err := Codecs[AttributeSigned].Encode(int32(-123456))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte{255, 254, 29, 192}) {
		t.Fail()
	}
	integer, err := Codecs[AttributeSigned].Decode([]byte{255, 255, 255, 255})
	if err != nil {
		t.Fatal(err)
	}
	if integer.(int32) != -1 {
		t.Fail()
	}
	ai := attributeSigned{}
	if v, _ := ai.Transform(int64(-1)); v.(int32) != -1 {
		t.Fail()
	}
	if v, _ := ai.Transform(uint(12)); v.(int32) != 12 {
		t.Fail()
	}
}

func TestAttributeSigned_Error(t *testing.T) {
	_, err := Codecs[AttributeInteger64].Encode(uint32(123456))
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeInteger64].Decode([]byte{255, 255, 255})
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeOctets(t *testing.T) {
	wired, err := Codecs[AttributeOctets].Encode("test")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte("test")) {
		t.Fail()
	}
	data, err := Codecs[AttributeOctets].Decode([]byte{255, 255, 255, 0, 255})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data.([]byte), []byte{255, 255, 255, 0, 255}) {
		t.Fail()
	}
}

func TestAttributeOctets_Error(t *testing.T) {
	_, err := Codecs[AttributeOctets].Encode(123456)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeString(t *testing.T) {
	wired, err := Codecs[AttributeString].Encode("test")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte("test")) {
		t.Fail()
	}
	wired, err = Codecs[AttributeString].Encode([]byte("test2"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte("test2")) {
		t.Fail()
	}
	data, err := Codecs[AttributeString].Decode([]byte("string"))
	if err != nil {
		t.Fatal(err)
	}
	if data.(string) != "string" {
		t.Fail()
	}
}

func TestAttributeString_Error(t *testing.T) {
	_, err := Codecs[AttributeString].Decode([]byte{255, 0, 0, 255})
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeAddress(t *testing.T) {
	wired, err := Codecs[AttributeIPAddr].Encode(net.ParseIP("10.3.8.213"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte{10, 3, 8, 213}) {
		t.Fail()
	}
	data, err := Codecs[AttributeIPAddr].Decode([]byte{1, 2, 3, 4})
	if err != nil {
		t.Fatal(err)
	}
	if data.(net.IP).String() != "1.2.3.4" {
		t.Fail()
	}
}

func TestAttributeAddress_Error(t *testing.T) {
	_, err := Codecs[AttributeIPAddr].Encode([]byte{0, 0, 255})
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeIPAddr].Encode(net.ParseIP("2001::1"))
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeIPAddr].Decode([]byte{0, 0, 255})
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeAddress6(t *testing.T) {
	wired, err := Codecs[AttributeIPv6Addr].Encode(net.ParseIP("2001::1"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte{32, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}) {
		t.Fail()
	}
	data, err := Codecs[AttributeIPv6Addr].Decode([]byte{0x20, 0x01, 0x0d, 0xa8, 0x02, 0x15, 0x8f, 0x02, 0x4c, 0xc1, 0xf8, 0x72, 0xd6, 0xd0, 0x8e, 0x5c})
	if err != nil {
		t.Fatal(err)
	}
	if data.(net.IP).String() != "2001:da8:215:8f02:4cc1:f872:d6d0:8e5c" {
		t.Fail()
	}
}

func TestAttributeAddress6_Error(t *testing.T) {
	_, err := Codecs[AttributeIPv6Addr].Encode([]byte{0, 0, 255})
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeIPv6Addr].Decode([]byte{0, 0, 255})
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeDate(t *testing.T) {
	wired, err := Codecs[AttributeDate].Encode(time.Unix(1514026068, 861))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte{90, 62, 52, 84}) {
		t.Fail()
	}
	data, err := Codecs[AttributeDate].Decode([]byte{0x5a, 0x3e, 0x32, 0x95})
	if err != nil {
		t.Fatal(err)
	}
	if data.(time.Time).String() != "2017-12-23 18:40:21 +0800 CST" {
		t.Fail()
	}
}

func TestAttributeDate_Error(t *testing.T) {
	_, err := Codecs[AttributeDate].Encode([]byte{0, 0, 255})
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeDate].Decode([]byte{0, 0, 255})
	if err == nil {
		t.Fatal(err)
	}
}

func TestAttributeShort(t *testing.T) {
	wired, err := Codecs[AttributeShort].Encode(uint16(1412))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wired, []byte{5, 132}) {
		t.Fail()
	}
	integer, err := Codecs[AttributeShort].Decode([]byte{255, 255})
	if err != nil {
		t.Fatal(err)
	}
	if integer.(uint16) != 65535 {
		t.Fail()
	}
	ai := attributeShort{}
	if v, _ := ai.Transform(int64(6000)); v.(uint16) != 6000 {
		t.Fail()
	}
	if v, _ := ai.Transform(uint64(6000)); v.(uint16) != 6000 {
		t.Fail()
	}
}

func TestAttributeShort_Error(t *testing.T) {
	_, err := Codecs[AttributeShort].Encode(uint32(123456))
	if err == nil {
		t.Fatal(err)
	}
	_, err = Codecs[AttributeShort].Decode([]byte{255, 255, 255})
	if err == nil {
		t.Fatal(err)
	}
}
