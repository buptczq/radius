package radius

import (
	"errors"
	"sync"
)

var builtinOnce sync.Once

// Builtin is the built-in dictionary. It is initially loaded with the
// attributes defined in RFC 2865 and RFC 2866.
var Builtin *Dictionary

func initDictionary() {
	Builtin = &Dictionary{}
}

type dictEntry struct {
	Type  byte
	OID   uint32
	Name  string
	Codec AttributeCodec
}

// Dictionary stores mappings between attribute names and types and
// AttributeCodecs.
type Dictionary struct {
	mu sync.RWMutex
	// OID, Type, Entry
	attributesByType map[uint32]map[byte]*dictEntry
	attributesByName map[string]*dictEntry
}

// Register registers the AttributeCodec for the given attribute name and type.
func (d *Dictionary) RegisterVendor(name string, oid uint32, t byte, codec AttributeType) error {
	d.mu.Lock()
	if d.attributesByType == nil {
		d.attributesByType = make(map[uint32]map[byte]*dictEntry)
	}
	if d.attributesByType[oid] == nil {
		d.attributesByType[oid] = make(map[byte]*dictEntry)
	}
	if d.attributesByType[oid][t] != nil {
		d.mu.Unlock()
		return errors.New("radius: attribute already registered")
	}
	entry := &dictEntry{
		Type:  t,
		OID:   oid,
		Name:  name,
		Codec: Codecs[codec],
	}
	d.attributesByType[oid][t] = entry
	if d.attributesByName == nil {
		d.attributesByName = make(map[string]*dictEntry)
	}
	d.attributesByName[name] = entry
	d.mu.Unlock()
	return nil
}

// Register registers the AttributeCodec for the given attribute name and type.
func (d *Dictionary) Register(name string, t byte, codec AttributeType) error {
	return d.RegisterVendor(name, 0, t, codec)
}

// MustRegister is a helper for Register that panics if it returns an error.
func (d *Dictionary) MustRegisterVendor(name string, oid uint32, t byte, codec AttributeType) {
	if err := d.RegisterVendor(name, oid, t, codec); err != nil {
		panic(err)
	}
}

// MustRegister is a helper for Register that panics if it returns an error.
func (d *Dictionary) MustRegister(name string, t byte, codec AttributeType) {
	d.MustRegisterVendor(name, 0, t, codec)
}

func (d *Dictionary) get(name string) (value *dictEntry, ok bool) {
	d.mu.RLock()
	entry := d.attributesByName[name]
	d.mu.RUnlock()
	if entry == nil {
		return
	}
	value = entry
	ok = true
	return
}

// Name returns the registered name for the given attribute type. ok is false
// if the given type is not registered.
func (d *Dictionary) Name(oid uint32, t byte) (name string, ok bool) {
	d.mu.RLock()
	entry := d.attributesByType[oid][t]
	d.mu.RUnlock()
	if entry == nil {
		return
	}
	name = entry.Name
	ok = true
	return
}

// Type returns the registered type for the given attribute name. ok is false
// if the given name is not registered.
func (d *Dictionary) Type(name string) (oid uint32, t byte, ok bool) {
	d.mu.RLock()
	entry := d.attributesByName[name]
	d.mu.RUnlock()
	if entry == nil {
		return
	}
	oid = entry.OID
	t = entry.Type
	ok = true
	return
}

// Codec returns the AttributeCodec for the given registered type. nil is
// returned if the given type is not registered.
func (d *Dictionary) Codec(oid uint32, t byte) AttributeCodec {
	d.mu.RLock()
	entry := d.attributesByType[oid][t]
	d.mu.RUnlock()
	if entry == nil {
		return Codecs[AttributeUnknown]
	}
	return entry.Codec
}

// Codec returns the AttributeCodec for the given registered type. nil is
// returned if the given type is not registered.
func (d *Dictionary) CodecByName(name string) AttributeCodec {
	d.mu.RLock()
	entry := d.attributesByName[name]
	d.mu.RUnlock()
	if entry == nil {
		return Codecs[AttributeUnknown]
	}
	return entry.Codec
}
