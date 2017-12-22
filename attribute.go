package radius

// Attribute is a RADIUS attribute, which is part of a RADIUS packet.
type Attribute []byte

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
