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
