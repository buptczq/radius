package radius

import (
	"errors"
	"fmt"
)

// TODO : support tag
// TODO : support encrypt

// If multiple Attributes with the same Type are present, the order of
// Attributes with the same Type MUST be preserved by any proxies.  The
// order of Attributes of different Types is not required to be
// preserved.
type Attributes map[byte][]Attribute

const VENDOR_SPECIFIC byte = 26

// ParseAttributes parses the wire-encoded RADIUS attributes and returns a new
// Attributes value. An error is returned if the buffer is malformed.
func ParseAttributes(b []byte) (Attributes, error) {
	attrs := make(map[byte][]Attribute)

	for len(b) > 0 {
		if len(b) < 2 {
			return nil, errors.New("attribute must be at least 2 bytes long")
		}

		attrLength := b[1]
		if attrLength < 1 || attrLength > 253 || len(b) < int(attrLength) {
			return nil, errors.New("invalid attribute length")
		}

		attrType := b[0]
		attrValue := b[2:attrLength]

		attrs[attrType] = append(attrs[attrType], attrValue)
		b = b[attrLength:]
	}

	return attrs, nil
}

// Add appends the given Attribute to the map entry of the given type.
func (a Attributes) AddRaw(key AttributeKey, value Attribute) {
	oid := key.Vendor()
	if oid == 0 {
		a[key.Type()] = append(a[key.Type()], value)
	} else {
		a[VENDOR_SPECIFIC] = append(a[VENDOR_SPECIFIC], EncodeAVPairByte(oid, key.Type(), value))
	}
}

// Del removes all Attributes of the given type from a.
func (a Attributes) Del(key AttributeKey) {
	oid := key.Vendor()
	if oid == 0 {
		delete(a, key.Type())
	} else {
		old := a[VENDOR_SPECIFIC]
		if old == nil {
			return
		}
		newAttrs := make([]Attribute, 0)
		for _, vsa := range old {
			oid, t, _, _ := DecodeAVPairByte(vsa)
			if MakeAttributeKey(oid, 0, t) != key {
				newAttrs = append(newAttrs, vsa)
			}
		}
		a[VENDOR_SPECIFIC] = newAttrs
	}
}

// Lookup returns the first Attribute of Type key. nil and false is returned if
// no Attribute of Type key exists in a.
func (a Attributes) LookupRaw(key AttributeKey) (Attribute, bool) {
	oid := key.Vendor()
	if oid == 0 {
		m := a[key.Type()]
		if len(m) == 0 {
			return nil, false
		}
		return m[0], true
	} else {
		attrs := a[VENDOR_SPECIFIC]
		if attrs == nil {
			return nil, false
		}
		for _, vsa := range attrs {
			oid, t, data, err := DecodeAVPairByte(vsa)
			if err != nil {
				return nil, false
			}
			if MakeAttributeKey(oid, 0, t) == key {
				return data, true
			}
		}
	}
	return nil, false
}

// Get returns the first Attribute of Type key. nil is returned if no Attribute
// of Type key exists in a.
func (a Attributes) GetRaw(key AttributeKey) Attribute {
	attr, _ := a.LookupRaw(key)
	return attr
}

// Set removes all Attributes of Type key and appends value.
func (a Attributes) SetRaw(key AttributeKey, value Attribute) {
	oid := key.Vendor()
	if oid == 0 {
		a[key.Type()] = []Attribute{value}
	} else {
		a.Del(key)
		a.AddRaw(key, value)
	}
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

// Set sets the value of the first attribute whose type matches the
// given key. If no such attribute exists, a new attribute is added
func (a Attributes) Set(key AttributeKey, value interface{}) error {
	entry, ok := Builtin.Get(key)
	if !ok {
		return fmt.Errorf("can't find attribute %s", key)
	}
	newValue := value
	if transformer, ok := entry.Codec().(AttributeTransformer); ok {
		transformed, err := transformer.Transform(value)
		if err != nil {
			return err
		}
		newValue = transformed
	}
	attr, err := entry.Codec().Encode(newValue)
	if err != nil {
		return err
	}
	a.SetRaw(entry.Key, attr)
	return nil
}

// SetByName sets the value of the first attribute whose dictionary name matches the
// given name. If no such attribute exists, a new attribute is added
func (a Attributes) SetByName(name string, value interface{}) error {
	entry, ok := Builtin.GetByName(name)
	if !ok {
		return fmt.Errorf("can't find attribute %s", name)
	}
	newValue := value
	if transformer, ok := entry.Codec().(AttributeTransformer); ok {
		transformed, err := transformer.Transform(value)
		if err != nil {
			return err
		}
		newValue = transformed
	}
	attr, err := entry.Codec().Encode(newValue)
	if err != nil {
		return err
	}
	a.SetRaw(entry.Key, attr)
	return nil
}

// Get returns the value of the first attribute whose type matches
// the given type. nil is returned if no such attribute exists.
func (a Attributes) Get(key AttributeKey) interface{} {
	entry, ok := Builtin.Get(key)
	if !ok {
		return fmt.Errorf("can't find attribute %s", key)
	}
	if data, ok := a.LookupRaw(key); ok {
		if v, err := entry.Codec().Decode(data); err == nil {
			return v
		}
	}
	return nil
}

// GetByName returns the value of the first attribute whose dictionary name matches
// the given name. nil is returned if no such attribute exists.
func (a Attributes) GetByName(name string) interface{} {
	entry, ok := Builtin.GetByName(name)
	if !ok {
		return nil
	}
	attr := a.GetRaw(entry.Key)
	if v, err := entry.Codec().Decode(attr); err == nil {
		return v
	}
	return nil
}

func (a Attributes) GetString(key AttributeKey) string {
	value := a.Get(key)
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
