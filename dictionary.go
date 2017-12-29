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

const (
	TAG_NONE    = 0
	TAG_INTERGE = 1
	TAG_STRING  = 2
)

// Dictionary stores mappings between attribute names and types and
// AttributeCodecs.
type Dictionary struct {
	mu               sync.RWMutex
	attributesByType map[AttributeKey]*DictEntry
	attributesByName map[string]*DictEntry
	attributesHasTag map[AttributeKey]int
}

// RegisterEx registers the AttributeCodecs for the given attribute options.
func (d *Dictionary) Register(name string, oid uint32, t byte, hasTag bool, encrypt int, attrType AttributeType) error {
	if encrypt > 3 || encrypt < 0 {
		return errors.New("illegal attribute encryption")
	}
	if hasTag && !(attrType == AttributeString || attrType == AttributeInteger) {
		return errors.New("only string and integer can have tags")
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
	if d.attributesHasTag == nil {
		d.attributesHasTag = make(map[AttributeKey]int)
	}
	if hasTag {
		if attrType == AttributeInteger {
			d.attributesHasTag[key] = TAG_INTERGE
		} else if attrType == AttributeString {
			d.attributesHasTag[key] = TAG_STRING
		}
	}
	d.mu.Unlock()
	return nil
}

// MustRegisterEx is a helper for Register that panics if it returns an error.
func (d *Dictionary) MustRegister(name string, oid uint32, code byte, hasTag bool, encrypt int, attrType AttributeType) {
	if err := d.Register(name, oid, code, hasTag, encrypt, attrType); err != nil {
		panic(err)
	}
}

// Get returns the registered entry for the given attribute key. ok is false
// if the given key is not registered.
func (d *Dictionary) Get(key AttributeKey) (value *DictEntry, ok bool) {
	d.mu.RLock()
	entry := d.attributesByType[key.WithoutTag()]
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

// Name returns the registered name for the given attribute key. ok is false
// if the given type is not registered.
func (d *Dictionary) Name(key AttributeKey) (name string, ok bool) {
	d.mu.RLock()
	entry := d.attributesByType[key.WithoutTag()]
	d.mu.RUnlock()
	if entry == nil {
		return
	}
	name = entry.Name
	ok = true
	return
}

// HasTag returns whether the given attribute's type has a tag
func (d *Dictionary) TagType(key AttributeKey) int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if v, ok := d.attributesHasTag[key.WithoutTag()]; ok {
		return v
	}
	return TAG_NONE
}

// Codecs returns the codecs for the given attribute.
func (d *DictEntry) Codecs() AttributeCodec {
	if codecs, ok := Codecs[d.AttrType]; ok {
		return codecs
	} else {
		return Codecs[AttributeUnknown]
	}
}
