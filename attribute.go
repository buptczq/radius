package radius

import "fmt"

// Attribute is a RADIUS attribute, which is part of a RADIUS packet.
type Attribute []byte

// AttributeKey represents a attribute key
// AttributeKey = VENDOR << 32 | TAG << 8 | Type
// TAG MUST be zero in dictionary
type AttributeKey uint64

// golang 1.9
//type AttrVendorId = uint32
//type AttrTag = byte
//type AttrType = byte

const (
	VENDOR_MASK uint64 = 0xFFFFFFFF << 32
	TAG_MASK    uint64 = 0xFF << 8
	TYPE_MASK   uint64 = 0xFF
)

func (k AttributeKey) Vendor() uint32 {
	return uint32((uint64(k) & VENDOR_MASK) >> 32)
}

func (k AttributeKey) Tag() byte {
	return byte((uint64(k) & TAG_MASK) >> 8)
}

func (k AttributeKey) Type() byte {
	return byte(uint64(k) & TYPE_MASK)
}

func (k AttributeKey) WithTag(tag byte) AttributeKey {
	return AttributeKey((uint64(k) & ^TAG_MASK) | uint64(tag)<<8)
}

func (k AttributeKey) WithoutTag() AttributeKey {
	return AttributeKey(uint64(k) & ^TAG_MASK)
}

func (k AttributeKey) ValidTag() bool {
	//The Tag field is one octet in length and is intended to provide a
	//means of grouping attributes in the same packet which refer to the
	//same tunnel.  Valid values for this field are 0x01 through 0x1F,
	//inclusive.  If the Tag field is unused, it MUST be zero (0x00).
	return k.Tag() > 0 && k.Tag() < 0x20
}

func (k AttributeKey) String() string {
	if name, ok := Builtin.Name(k); ok {
		if Builtin.TagType(k) > 0 {
			return fmt.Sprintf("%s:%d", name, k.Tag())
		} else {
			return name
		}
	} else {
		return fmt.Sprintf("(Vendor = %d, Type = %d, Tag = %d)", k.Vendor(), k.Type(), k.Tag())
	}
}

func MakeAttributeKey(oid uint32, tag byte, attrType byte) AttributeKey {
	return AttributeKey(uint64(oid)<<32 | uint64(tag)<<8 | uint64(attrType))
}

// AttributeCodec defines how an Attribute is encoded and decoded to and from
// wire data.
type AttributeCodec interface {
	// Note: do not store wire; make a copy of it.
	Decode(wire Attribute) (interface{}, error)
	Encode(value interface{}) (Attribute, error)
}

// AttributeTransformer defines an extension of AttributeCodec. It provides a
// method for converting attribute values to ones permitted by the attribute.
type AttributeTransformer interface {
	Transform(value interface{}) (interface{}, error)
}

// AttributeStringer defines an extension of AttributeCodec. It provides a
// method for converting an attribute value to a string.
type AttributeStringer interface {
	String(value interface{}) string
}
