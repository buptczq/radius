package radius

import (
	"bytes"
	"errors"
	"fmt"
)

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

// AddRaw appends the given Attribute to the map entry of the given key.
func (a Attributes) AddRaw(key AttributeKey, value Attribute) {
	oid := key.Vendor()
	tagType := Builtin.TagType(key)
	if oid == 0 {
		switch tagType {
		case TAG_NONE:
			a[key.Type()] = append(a[key.Type()], value)
		case TAG_STRING:
			if key.ValidTag() {
				a[key.Type()] = append(a[key.Type()], AddTag(value, key.Tag()))
			} else {
				a[key.Type()] = append(a[key.Type()], value)
			}
		case TAG_INTERGE:
			if key.ValidTag() {
				a[key.Type()] = append(a[key.Type()], AddTag(value, key.Tag()))
			} else {
				a[key.Type()] = append(a[key.Type()], AddTag(value, 0))
			}
		}
	} else {
		switch tagType {
		case TAG_NONE:
			a[VENDOR_SPECIFIC] = append(a[VENDOR_SPECIFIC], EncodeAVPairByte(oid, key.Type(), value))
		case TAG_STRING:
			if key.ValidTag() {
				a[VENDOR_SPECIFIC] = append(a[VENDOR_SPECIFIC], EncodeAVPairByteTag(oid, key.Type(), key.Tag(), value))
			} else {
				a[VENDOR_SPECIFIC] = append(a[VENDOR_SPECIFIC], EncodeAVPairByte(oid, key.Type(), value))
			}
		case TAG_INTERGE:
			if key.ValidTag() {
				a[VENDOR_SPECIFIC] = append(a[VENDOR_SPECIFIC], EncodeAVPairByteTag(oid, key.Type(), key.Tag(), value))
			} else {
				a[VENDOR_SPECIFIC] = append(a[VENDOR_SPECIFIC], EncodeAVPairByteTag(oid, key.Type(), 0, value))
			}
		}
	}
}

// Del removes all Attributes of the given key from a.
func (a Attributes) Del(key AttributeKey) {
	oid := key.Vendor()
	hasTag := Builtin.TagType(key) > 0
	if oid == 0 {
		if hasTag && key.Tag() != 0 {
			old := a[key.Type()]
			if old == nil {
				return
			}
			newAttrs := make([]Attribute, 0)
			for _, data := range old {
				if len(data) == 0 || data[0] != key.Tag() {
					newAttrs = append(newAttrs, data)
				}
			}
			a[key.Type()] = newAttrs
		} else {
			// tag == 0, delete all attributes
			delete(a, key.Type())
		}
	} else {
		old := a[VENDOR_SPECIFIC]
		if old == nil {
			return
		}
		newAttrs := make([]Attribute, 0)
		for _, vsa := range old {
			oid, t, data, err := DecodeAVPairByte(vsa)
			if err != nil {
				// ignore malformation
				continue
			}
			if hasTag && key.Tag() != 0 {
				if len(data) > 0 {
					if MakeAttributeKey(oid, data[0], t) != key {
						newAttrs = append(newAttrs, vsa)
					}
				}
			} else {
				// ignore tag
				if MakeAttributeKey(oid, 0, t) != key.WithoutTag() {
					newAttrs = append(newAttrs, vsa)
				}
			}
		}
		a[VENDOR_SPECIFIC] = newAttrs
	}
}

func stripTag(tagType int, data []byte) ([]byte, bool) {
	switch tagType {
	case TAG_NONE:
		return data, true
	case TAG_INTERGE:
		if len(data) > 0 {
			return data[1:], true
		}
		return nil, false
	case TAG_STRING:
		if len(data) > 0 {
			if data[0] < 0x20 {
				return data[1:], true
			} else {
				return data, true
			}
		}
		return data, true
	}
	return nil, false
}

// LookupRaw returns the first Attribute of the given key. nil and false is returned if
// no Attribute of Type key exists in a. The tag field will be striped if has.
func (a Attributes) LookupRaw(key AttributeKey) (Attribute, bool) {
	oid := key.Vendor()
	tagType := Builtin.TagType(key)
	if oid == 0 {
		m := a[key.Type()]
		if len(m) == 0 {
			return nil, false
		}
		if tagType > 0 && key.ValidTag() {
			for _, data := range m {
				if len(data) > 0 && data[0] == key.Tag() {
					return data[1:], true
				}
			}
		} else {
			return stripTag(tagType, m[0])
		}
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
			if tagType > 0 && key.ValidTag() {
				if len(data) > 0 && data[0] == key.Tag() {
					return data[1:], true
				}
			} else {
				// ignore tag
				if MakeAttributeKey(oid, 0, t) == key.WithoutTag() {
					return stripTag(tagType, data)
				}
			}
		}
	}
	return nil, false
}

// GetRaw returns the first Attribute of the given key. nil is returned if no Attribute
// of key exists in a. The tag field will be striped if has.
func (a Attributes) GetRaw(key AttributeKey) Attribute {
	attr, _ := a.LookupRaw(key)
	return attr
}

// SetRaw removes all Attributes of key and appends value.
func (a Attributes) SetRaw(key AttributeKey, value Attribute) {
	a.Del(key)
	a.AddRaw(key, value)
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
	if transformer, ok := entry.Codecs().(AttributeTransformer); ok {
		transformed, err := transformer.Transform(value)
		if err != nil {
			return err
		}
		newValue = transformed
	}
	attr, err := entry.Codecs().Encode(newValue)
	if err != nil {
		return err
	}
	a.SetRaw(key, attr)
	return nil
}

// SetByName sets the value of the first attribute whose dictionary name matches the
// given name. If no such attribute exists, a new attribute is added
func (a Attributes) SetByName(name string, value interface{}) error {
	return a.SetByNameWithTag(name, 0, value)
}

// SetByNameWithTag sets the value of the first attribute whose dictionary name matches the
// given name with a tag. If no such attribute exists, a new attribute is added
func (a Attributes) SetByNameWithTag(name string, tag byte, value interface{}) error {
	entry, ok := Builtin.GetByName(name)
	if !ok {
		return fmt.Errorf("can't find attribute %s", name)
	}
	newValue := value
	if transformer, ok := entry.Codecs().(AttributeTransformer); ok {
		transformed, err := transformer.Transform(value)
		if err != nil {
			return err
		}
		newValue = transformed
	}
	attr, err := entry.Codecs().Encode(newValue)
	if err != nil {
		return err
	}
	a.SetRaw(entry.Key.WithTag(tag), attr)
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
		if v, err := entry.Codecs().Decode(data); err == nil {
			return v
		}
	}
	return nil
}

// GetByNameWithTag returns the value of the first attribute whose dictionary name matches
// the given name and tag. nil is returned if no such attribute exists.
func (a Attributes) GetByNameWithTag(name string, tag byte) interface{} {
	entry, ok := Builtin.GetByName(name)
	if !ok {
		return nil
	}
	return a.Get(entry.Key.WithTag(tag))
}

// GetByName returns the value of the first attribute whose dictionary name matches
// the given name. nil is returned if no such attribute exists.
func (a Attributes) GetByName(name string) interface{} {
	return a.GetByNameWithTag(name, 0)
}

// GetString is a helper for getting a formatted string from the given attribute.
func (a Attributes) GetString(key AttributeKey) string {
	value := a.Get(key)
	if value == nil {
		return ""
	}
	switch v := value.(type) {
	case []byte:
		// FreeRADIUS style
		var b bytes.Buffer
		if len(v) == 0 {
			return ""
		}
		b.WriteString("0x")
		for _, x := range v {
			b.WriteString(fmt.Sprintf("%02X", x))
		}
		return b.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
