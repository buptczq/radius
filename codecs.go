package radius

import (
	"encoding/binary"
	"errors"
	"net"
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
	Codecs[AttributeString] = attributeText{}
	Codecs[AttributeOctets] = attributeOctets{}
	Codecs[AttributeIPAddr] = attributeAddress{}
	Codecs[AttributeDate] = attributeTime{}
	Codecs[AttributeInteger] = attributeInteger{}
	Codecs[AttributeIPv6Addr] = attributeAddress6{}
	// TODO: ipv6 prefix
	Codecs[AttributeIPv6Prefix] = attributeOctets{}
	Codecs[AttributeIFID] = attributeOctets{}
	Codecs[AttributeInteger64] = attributeInteger64{}
	Codecs[AttributeVSA] = attributeOctets{}
	Codecs[AttributeUnknown] = attributeOctets{}
}

type attributeText struct{}

func (attributeText) Decode(value Attribute) (interface{}, error) {
	if !utf8.Valid(value) {
		return nil, errors.New("radius: text attribute is not valid UTF-8")
	}
	return string(value), nil
}

func (attributeText) Encode(value interface{}) (Attribute, error) {
	str, ok := value.(string)
	if ok {
		return Attribute(str), nil
	}
	raw, ok := value.([]byte)
	if ok {
		return raw, nil
	}
	return nil, errors.New("radius: text attribute must be string or []byte")
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
	return nil, errors.New("radius: string attribute must be Attribute or string")
}

type attributeAddress struct{}

func (attributeAddress) Decode(value Attribute) (interface{}, error) {
	if len(value) != net.IPv4len {
		return nil, errors.New("radius: address attribute has invalid size")
	}
	v := make(Attribute, len(value))
	copy(v, value)
	return net.IP(v), nil
}

func (attributeAddress) Encode(value interface{}) (Attribute, error) {
	ip, ok := value.(net.IP)
	if !ok {
		return nil, errors.New("radius: address attribute must be net.IP")
	}
	ip = ip.To4()
	if ip == nil {
		return nil, errors.New("radius: address attribute must be an IPv4 net.IP")
	}
	return Attribute(ip), nil
}

type attributeAddress6 struct{}

func (attributeAddress6) Decode(value Attribute) (interface{}, error) {
	if len(value) != net.IPv6len {
		return nil, errors.New("radius: address attribute has invalid size")
	}
	v := make(Attribute, len(value))
	copy(v, value)
	return net.IP(v), nil
}

func (attributeAddress6) Encode(value interface{}) (Attribute, error) {
	ip, ok := value.(net.IP)
	if !ok {
		return nil, errors.New("radius: address attribute must be net.IP")
	}
	ip = ip.To16()
	if ip == nil {
		return nil, errors.New("radius: address attribute must be an IPv6 net.IP")
	}
	return Attribute(ip), nil
}

type attributeSigned struct{}

func (attributeSigned) Decode(value Attribute) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("radius: signed attribute has invalid size")
	}
	return int32(binary.BigEndian.Uint32(value)), nil
}

func (attributeSigned) Encode(value interface{}) (Attribute, error) {
	integer, ok := value.(int32)
	if !ok {
		return nil, errors.New("radius: signed attribute must be int32")
	}
	raw := make(Attribute, 4)
	binary.BigEndian.PutUint32(raw, (uint32)(integer))
	return raw, nil
}

type attributeInteger struct{}

func (attributeInteger) Decode(value Attribute) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("radius: integer attribute has invalid size")
	}
	return binary.BigEndian.Uint32(value), nil
}

func (attributeInteger) Encode(value interface{}) (Attribute, error) {
	integer, ok := value.(uint32)
	if !ok {
		return nil, errors.New("radius: integer attribute must be uint32")
	}
	raw := make(Attribute, 4)
	binary.BigEndian.PutUint32(raw, integer)
	return raw, nil
}

type attributeInteger64 struct{}

func (attributeInteger64) Decode(value Attribute) (interface{}, error) {
	if len(value) != 8 {
		return nil, errors.New("radius: integer64 attribute has invalid size")
	}
	return binary.BigEndian.Uint64(value), nil
}

func (attributeInteger64) Encode(value interface{}) (Attribute, error) {
	integer, ok := value.(uint64)
	if !ok {
		return nil, errors.New("radius: integer64 attribute must be uint64")
	}
	raw := make(Attribute, 8)
	binary.BigEndian.PutUint64(raw, integer)
	return raw, nil
}

type attributeTime struct{}

func (attributeTime) Decode(value Attribute) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("radius: time attribute has invalid size")
	}
	return time.Unix(int64(binary.BigEndian.Uint32(value)), 0), nil
}

func (attributeTime) Encode(value interface{}) (Attribute, error) {
	timestamp, ok := value.(time.Time)
	if !ok {
		return nil, errors.New("radius: time attribute must be time.Time")
	}
	raw := make(Attribute, 4)
	binary.BigEndian.PutUint32(raw, uint32(timestamp.Unix()))
	return raw, nil
}