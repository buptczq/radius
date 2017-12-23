package radius

import (
	"fmt"
	"testing"
)

func TestMakeAttributeKey(t *testing.T) {
	key := MakeAttributeKey(334324, 9, 16)
	if key.Type() != 16 {
		t.Fail()
	}
	if key.Vendor() != 334324 {
		t.Fail()
	}
	if key.Tag() != 9 {
		t.Fail()
	}
	if fmt.Sprint(key) != "(Vendor = 334324, Type = 16, Tag = 9)" {
		t.Fatal(key)
	}
}

func TestAttributeKey_WithTag(t *testing.T) {
	key := MakeAttributeKey(334324, 9, 16)
	newKey := key.WithTag(8)
	if newKey.Tag() != 8 || newKey.Type() != 16 || newKey.Vendor() != 334324 {
		t.Fail()
	}
}

func TestAttributeKey_WithoutTagTag(t *testing.T) {
	key := MakeAttributeKey(334324, 9, 16)
	newKey := key.WithoutTag()
	if newKey.Tag() != 0 || newKey.Type() != 16 || newKey.Vendor() != 334324 {
		t.Fail()
	}
}
