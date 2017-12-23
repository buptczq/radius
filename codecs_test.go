package radius

import (
	"bytes"
	"testing"
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
