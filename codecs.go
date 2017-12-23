package radius

import (
	"encoding/binary"
	"errors"
	"net"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)

type AttributeType int

const (
	AttributeString AttributeType = iota + 1
	AttributeOctets
	AttributeIPAddr
	AttributeDate
	AttributeInteger
	AttributeSigned
	AttributeIPv6Addr
	AttributeIPv6Prefix
	AttributeIFID
	AttributeInteger64

	AttributeVSA
	AttributeUnknown = -1
	// TODO: non-standard types?
)

func (t AttributeType) String() string {
	switch t {
	case AttributeString:
		return "string"
	case AttributeOctets:
		return "octets"
	case AttributeIPAddr:
		return "ipaddr"
	case AttributeDate:
		return "date"
	case AttributeInteger:
		return "integer"
	case AttributeSigned:
		return "signed"
	case AttributeIPv6Addr:
		return "ipv6addr"
	case AttributeIPv6Prefix:
		return "ipv6prefix"
	case AttributeIFID:
		return "ifid"
	case AttributeInteger64:
		return "integer64"
	case AttributeVSA:
		return "vsa"
	case AttributeUnknown:
		return "unknown"
	}
	return "AttributeType(" + strconv.Itoa(int(t)) + ")"
}

var Codecs map[AttributeType]AttributeCodec

func init() {
	Codecs = make(map[AttributeType]AttributeCodec)
	Codecs[AttributeString] = attributeString{}
	Codecs[AttributeOctets] = attributeOctets{}
	Codecs[AttributeIPAddr] = attributeAddress{}
	Codecs[AttributeDate] = attributeDate{}
	Codecs[AttributeInteger] = attributeInteger{}
	Codecs[AttributeSigned] = attributeSigned{}
	Codecs[AttributeIPv6Addr] = attributeAddress6{}
	// TODO: ipv6 prefix
	Codecs[AttributeIPv6Prefix] = attributeOctets{}
	Codecs[AttributeIFID] = attributeOctets{}
	Codecs[AttributeInteger64] = attributeInteger64{}
	Codecs[AttributeVSA] = attributeOctets{}
	Codecs[AttributeUnknown] = attributeOctets{}
}

type attributeString struct{}

func (attributeString) Decode(value Attribute) (interface{}, error) {
	if !utf8.Valid(value) {
		return nil, errors.New("text attribute is not valid UTF-8")
	}
	return string(value), nil
}

func (attributeString) Encode(value interface{}) (Attribute, error) {
	str, ok := value.(string)
	if ok {
		return Attribute(str), nil
	}
	raw, ok := value.([]byte)
	if ok {
		return raw, nil
	}
	return nil, errors.New("text attribute must be string or []byte")
}

type attributeOctets struct{}

func (attributeOctets) Decode(value Attribute) (interface{}, error) {
	v := make([]byte, len(value))
	copy(v, value)
	return v, nil
}

func (attributeOctets) Encode(value interface{}) (Attribute, error) {
	raw, ok := value.(Attribute)
	if ok {
		return raw, nil
	}
	str, ok := value.(string)
	if ok {
		return Attribute(str), nil
	}
	data, ok := value.([]byte)
	if ok {
		return Attribute(data), nil
	}
	return nil, errors.New("string attribute must be Attribute or string")
}

type attributeAddress struct{}

func (attributeAddress) Decode(value Attribute) (interface{}, error) {
	if len(value) != net.IPv4len {
		return nil, errors.New("address attribute has invalid size")
	}
	v := make(Attribute, len(value))
	copy(v, value)
	return net.IP(v), nil
}

func (attributeAddress) Encode(value interface{}) (Attribute, error) {
	ip, ok := value.(net.IP)
	if !ok {
		return nil, errors.New("address attribute must be net.IP")
	}
	ip = ip.To4()
	if ip == nil {
		return nil, errors.New("address attribute must be an IPv4 net.IP")
	}
	return Attribute(ip), nil
}

type attributeAddress6 struct{}

func (attributeAddress6) Decode(value Attribute) (interface{}, error) {
	if len(value) != net.IPv6len {
		return nil, errors.New("address attribute has invalid size")
	}
	v := make(Attribute, len(value))
	copy(v, value)
	return net.IP(v), nil
}

func (attributeAddress6) Encode(value interface{}) (Attribute, error) {
	ip, ok := value.(net.IP)
	if !ok {
		return nil, errors.New("address attribute must be net.IP")
	}
	ip = ip.To16()
	if ip == nil {
		return nil, errors.New("address attribute must be an IPv6 net.IP")
	}
	return Attribute(ip), nil
}

type attributeSigned struct{}

func (attributeSigned) Decode(value Attribute) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("signed attribute has invalid size")
	}
	return int32(binary.BigEndian.Uint32(value)), nil
}

func (attributeSigned) Encode(value interface{}) (Attribute, error) {
	integer, ok := value.(int32)
	if !ok {
		return nil, errors.New("signed attribute must be int32")
	}
	raw := make(Attribute, 4)
	binary.BigEndian.PutUint32(raw, (uint32)(integer))
	return raw, nil
}

func (attributeSigned) Transform(value interface{}) (interface{}, error) {
	switch t := value.(type) {
	case int, int8, int16, int32, int64:
		a := reflect.ValueOf(t).Int() // a has type int64
		return int32(a), nil
	case uint, uint8, uint16, uint32, uint64:
		a := reflect.ValueOf(t).Uint() // a has type uint64
		return int32(a), nil
	}
	return nil, errors.New("invalid integer attribute")
}

type attributeInteger struct{}

func (attributeInteger) Decode(value Attribute) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("integer attribute has invalid size")
	}
	return binary.BigEndian.Uint32(value), nil
}

func (attributeInteger) Encode(value interface{}) (Attribute, error) {
	integer, ok := value.(uint32)
	if !ok {
		return nil, errors.New("integer attribute must be uint32")
	}
	raw := make(Attribute, 4)
	binary.BigEndian.PutUint32(raw, integer)
	return raw, nil
}

func (attributeInteger) Transform(value interface{}) (interface{}, error) {
	switch t := value.(type) {
	case int, int8, int16, int32, int64:
		a := reflect.ValueOf(t).Int() // a has type int64
		return uint32(a), nil
	case uint, uint8, uint16, uint32, uint64:
		a := reflect.ValueOf(t).Uint() // a has type uint64
		return uint32(a), nil
	}
	return nil, errors.New("invalid integer attribute")
}

type attributeInteger64 struct{}

func (attributeInteger64) Decode(value Attribute) (interface{}, error) {
	if len(value) != 8 {
		return nil, errors.New("integer64 attribute has invalid size")
	}
	return binary.BigEndian.Uint64(value), nil
}

func (attributeInteger64) Encode(value interface{}) (Attribute, error) {
	integer, ok := value.(uint64)
	if !ok {
		return nil, errors.New("integer64 attribute must be uint64")
	}
	raw := make(Attribute, 8)
	binary.BigEndian.PutUint64(raw, integer)
	return raw, nil
}

func (attributeInteger64) Transform(value interface{}) (interface{}, error) {
	switch t := value.(type) {
	case int, int8, int16, int32, int64:
		a := reflect.ValueOf(t).Int() // a has type int64
		return uint64(a), nil
	case uint, uint8, uint16, uint32, uint64:
		a := reflect.ValueOf(t).Uint() // a has type uint64
		return uint64(a), nil
	}
	return nil, errors.New("invalid integer attribute")
}

type attributeDate struct{}

func (attributeDate) Decode(value Attribute) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("time attribute has invalid size")
	}
	return time.Unix(int64(binary.BigEndian.Uint32(value)), 0), nil
}

func (attributeDate) Encode(value interface{}) (Attribute, error) {
	timestamp, ok := value.(time.Time)
	if !ok {
		return nil, errors.New("time attribute must be time.Time")
	}
	raw := make(Attribute, 4)
	binary.BigEndian.PutUint32(raw, uint32(timestamp.Unix()))
	return raw, nil
}
