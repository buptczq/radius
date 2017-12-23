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

type DictEntry struct {
	Name     string
	Key      AttributeKey
	Encrypt  int
	HasTag   bool
	AttrType AttributeType
}

// Dictionary stores mappings between attribute names and types and
// AttributeCodecs.
type Dictionary struct {
	mu               sync.RWMutex
	attributesByType map[AttributeKey]*DictEntry
	attributesByName map[string]*DictEntry
}

// RegisterEx registers the AttributeCodec for the given attribute options.
func (d *Dictionary) RegisterEx(name string, oid uint32, t byte, hasTag bool, encrypt int, attrType AttributeType) error {
	if encrypt > 3 || encrypt < 0 {
		return errors.New("illegal attribute encryption")
	}
	key := MakeAttributeKey(oid, 0, t)
	d.mu.Lock()
	if d.attributesByType == nil {
		d.attributesByType = make(map[AttributeKey]*DictEntry)
	}
	if d.attributesByType[key] != nil {
		d.mu.Unlock()
		return errors.New("attribute already registered")
	}
	entry := &DictEntry{
		Name:     name,
		Key:      key,
		AttrType: attrType,
		HasTag:   hasTag,
		Encrypt:  encrypt,
	}
	d.attributesByType[key] = entry
	if d.attributesByName == nil {
		d.attributesByName = make(map[string]*DictEntry)
	}
	d.attributesByName[name] = entry
	d.mu.Unlock()
	return nil
}

// Register registers the AttributeCodec for the given attribute name and type.
func (d *Dictionary) Register(name string, code byte, attrType AttributeType) error {
	return d.RegisterEx(name, 0, code, false, 0, attrType)
}

// MustRegisterVendor is a helper for Register that panics if it returns an error.
func (d *Dictionary) MustRegisterVendor(name string, oid uint32, code byte, attrType AttributeType) {
	if err := d.RegisterEx(name, oid, code, false, 0, attrType); err != nil {
		panic(err)
	}
}

// MustRegisterEx is a helper for Register that panics if it returns an error.
func (d *Dictionary) MustRegisterEx(name string, oid uint32, code byte, hasTag bool, encrypt int, attrType AttributeType) {
	if err := d.RegisterEx(name, oid, code, hasTag, encrypt, attrType); err != nil {
		panic(err)
	}
}

// MustRegister is a helper for Register that panics if it returns an error.
func (d *Dictionary) MustRegister(name string, code byte, attrType AttributeType) {
	d.MustRegisterVendor(name, 0, code, attrType)
}

// Get returns the registered entry for the given attribute key. ok is false
// if the given key is not registered.
func (d *Dictionary) Get(key AttributeKey) (value *DictEntry, ok bool) {
	d.mu.RLock()
	entry := d.attributesByType[key]
	d.mu.RUnlock()
	if entry == nil {
		return
	}
	value = entry
	ok = true
	return
}

// GetByName returns the registered entry for the given attribute name. ok is false
// if the given name is not registered.
func (d *Dictionary) GetByName(name string) (value *DictEntry, ok bool) {
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
func (d *Dictionary) Name(key AttributeKey) (name string, ok bool) {
	d.mu.RLock()
	entry := d.attributesByType[key]
	d.mu.RUnlock()
	if entry == nil {
		return
	}
	name = entry.Name
	ok = true
	return
}

// Codec returns the codec for the given attribute.
func (d *DictEntry) Codec() AttributeCodec {
	if codec, ok := Codecs[d.AttrType]; ok {
		return codec
	} else {
		return Codecs[AttributeUnknown]
	}
}
