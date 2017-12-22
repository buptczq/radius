package radius

import (
	"errors"
	"fmt"
)

type Attributes map[byte][]Attribute

// ParseAttributes parses the wire-encoded RADIUS attributes and returns a new
// Attributes value. An error is returned if the buffer is malformed.
func ParseAttributes(b []byte) (Attributes, error) {
	attrs := make(map[byte][]Attribute)

	for len(b) > 0 {
		if len(b) < 2 {
			return nil, errors.New("radius: attribute must be at least 2 bytes long")
		}

		attrLength := b[1]
		if attrLength < 1 || attrLength > 253 || len(b) < int(attrLength) {
			return nil, errors.New("radius: invalid attribute length")
		}

		attrType := b[0]
		attrValue := b[2:attrLength]

		attrs[attrType] = append(attrs[attrType], attrValue)
		b = b[attrLength:]
	}

	return attrs, nil
}

// Add appends the given Attribute to the map entry of the given type.
func (a Attributes) Add(key byte, value Attribute) {
	a[key] = append(a[key], value)
}

// Del removes all Attributes of the given type from a.
func (a Attributes) Del(key byte) {
	delete(a, key)
}

// Get returns the first Attribute of Type key. nil is returned if no Attribute
// of Type key exists in a.
func (a Attributes) GetRaw(key byte) Attribute {
	attr, _ := a.Lookup(key)
	return attr
}

// Lookup returns the first Attribute of Type key. nil and false is returned if
// no Attribute of Type key exists in a.
func (a Attributes) Lookup(key byte) (Attribute, bool) {
	m := a[key]
	if len(m) == 0 {
		return nil, false
	}
	return m[0], true
}

// Set removes all Attributes of Type key and appends value.
func (a Attributes) SetAttr(key byte, value Attribute) {
	a[key] = []Attribute{value}
}

// Len returns the total number of Attributes in a.
func (a Attributes) Len() int {
	var i int
	for _, s := range a {
		i += len(s)
	}
	return i
}

func (a Attributes) encodeTo(b []byte) {
	for typ, attrs := range a {
		if typ < 0 || typ > 255 {
			continue
		}
		for _, attr := range attrs {
			size := 2 + len(attr)
			b[0] = byte(typ)
			b[1] = byte(size)
			copy(b[2:], attr)
			b = b[size:]
		}
	}
}

func (a Attributes) wireSize() (bytes int) {
	for typ, attrs := range a {
		if typ < 0 || typ > 255 {
			continue
		}
		for _, attr := range attrs {
			// type field + length field + value field
			bytes += 1 + 1 + len(attr)
		}
	}
	return
}

// Value returns the value of the first attribute whose dictionary name matches
// the given name. nil is returned if no such attribute exists.
func (a Attributes) Value(name string) interface{} {
	entry, ok := Builtin.get(name)
	if !ok {
		return nil
	}
	if entry.OID == 0 {
		if data, ok := a.Lookup(entry.Type); ok {
			if v, err := entry.Codec.Decode(data); err == nil {
				return v
			}
		}
	} else {
		m := a[26]
		if len(m) == 0 {
			return nil
		}
		for _, vsa := range m {
			oid, t, data, _ := DecodeAVPairByte(vsa)
			if oid == entry.OID && t == entry.Type {
				if v, err := entry.Codec.Decode(data); err == nil {
					return v
				}
			}
		}
	}
	return nil
}

// Set sets the value of the first attribute whose dictionary name matches the
// given name. If no such attribute exists, a new attribute is added
func (a Attributes) Set(name string, value interface{}) error {
	entry, ok := Builtin.get(name)
	if !ok {
		return fmt.Errorf("can't find attribute %s", entry)
	}
	newValue := value
	if transformer, ok := entry.Codec.(AttributeTransformer); ok {
		transformed, err := transformer.Transform(value)
		if err != nil {
			return err
		}
		newValue = transformed
	}
	attr, err := entry.Codec.Encode(newValue)
	if err != nil {
		return err
	}
	if entry.OID == 0 {
		a.SetAttr(entry.Type, attr)
	} else {
		a.Add(26, EncodeAVPairByte(entry.OID, entry.Type, attr))
	}
	return nil
}

// Get returns the value of the first attribute whose type matches
// the given type. nil is returned if no such attribute exists.
func (a Attributes) Get(t byte) interface{} {
	if data, ok := a.Lookup(t); ok {
		if v, err := Builtin.Codec(0, t).Decode(data); err == nil {
			return v
		}
	}
	return nil
}

func (a Attributes) GetString(t byte) string {
	if data, ok := a.Lookup(t); ok {
		if v, err := Builtin.Codec(0, t).Decode(data); err == nil {
			return fmt.Sprintf("%v", v)
		}
	}
	return ""
}
