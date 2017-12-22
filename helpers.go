package radius

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
)

// Encodes AVPair into Vendor-Specific attribute format with tag (string)
func EncodeAVpairTag(vendor_id uint32, type_id uint8, tag uint8, value string) (vsa []byte) {
	return EncodeAVPairByteTag(vendor_id, type_id, tag, []byte(value))
}

// Encodes AVPair into Vendor-Specific attribute format (byte)
func EncodeAVPairByte(vendor_id uint32, type_id uint8, value []byte) (vsa []byte) {
	var b bytes.Buffer
	bv := make([]byte, 4)
	binary.BigEndian.PutUint32(bv, vendor_id)

	// Vendor-Id(4) + Type-ID(1) + Length(1)
	b.Write(bv)
	b.Write([]byte{byte(type_id), byte(len(value) + 2)})

	// Append attribute value pair
	b.Write(value)

	vsa = b.Bytes()
	return
}

// Encodes AVPair into Vendor-Specific attribute format with tag (byte)
func EncodeAVPairByteTag(vendor_id uint32, type_id uint8, tag uint8, value []byte) (vsa []byte) {
	var b bytes.Buffer
	bv := make([]byte, 4)
	binary.BigEndian.PutUint32(bv, vendor_id)

	// Vendor-Id(4) + Type-ID(1) + Length(1)
	b.Write(bv)
	b.Write([]byte{byte(type_id), byte(len(value) + 3)})

	// Add tag
	b.WriteByte(byte(tag))

	// Append attribute value pair
	b.Write(value)

	vsa = b.Bytes()
	return
}

// Decodes VSA (byte)
func DecodeAVPairByte(vsa []byte) (vendor_id uint32, type_id uint8, value []byte, err error) {
	if len(vsa) <= 6 {
		err = fmt.Errorf("Too short VSA: %d bytes", len(vsa))
		return
	}

	vendor_id = binary.BigEndian.Uint32([]byte{vsa[0], vsa[1], vsa[2], vsa[3]})
	type_id = uint8(vsa[4])
	value = vsa[6:]
	return
}

// UserPassword decrypts the given  "User-Password"-encrypted (as defined in RFC
// 2865) Attribute, and returns the plaintext. An error is returned if the
// attribute length is invalid, the secret is empty, or the requestAuthenticator
// length is invalid.
func UserPassword(a Attribute, secret, requestAuthenticator []byte) ([]byte, error) {
	if len(a) < 16 || len(a) > 128 {
		return nil, errors.New("invalid attribute length (" + strconv.Itoa(len(a)) + ")")
	}
	if len(secret) == 0 {
		return nil, errors.New("empty secret")
	}
	if len(requestAuthenticator) != 16 {
		return nil, errors.New("invalid requestAuthenticator length (" + strconv.Itoa(len(requestAuthenticator)) + ")")
	}

	dec := make([]byte, 0, len(a))

	hash := md5.New()
	hash.Write(secret)
	hash.Write(requestAuthenticator)
	dec = hash.Sum(dec)

	for i, b := range a[:16] {
		dec[i] ^= b
	}

	for i := 16; i < len(a); i += 16 {
		hash.Reset()
		hash.Write(secret)
		hash.Write(a[i-16 : i])
		dec = hash.Sum(dec)

		for j, b := range a[i : i+16] {
			dec[i+j] ^= b
		}
	}

	if i := bytes.IndexByte(dec, 0); i > -1 {
		return dec[:i], nil
	}
	return dec, nil
}

// NewUserPassword returns a new "User-Password"-encrypted attribute from the
// given plaintext, secret, and requestAuthenticator. An error is returned if
// the plaintext is too long, the secret is empty, or the requestAuthenticator
// is an invalid length.
func NewUserPassword(plaintext, secret, requestAuthenticator []byte) (Attribute, error) {
	if len(plaintext) > 128 {
		return nil, errors.New("plaintext longer than 128 characters")
	}
	if len(secret) == 0 {
		return nil, errors.New("empty secret")
	}
	if len(requestAuthenticator) != 16 {
		return nil, errors.New("requestAuthenticator not 16-bytes")
	}

	chunks := len(plaintext) >> 4
	if chunks == 0 {
		chunks = 1
	}

	enc := make([]byte, 0, chunks*16)

	hash := md5.New()
	hash.Write(secret)
	hash.Write(requestAuthenticator)
	enc = hash.Sum(enc)

	for i, b := range plaintext[:16] {
		enc[i] ^= b
	}

	for i := 16; i < len(plaintext); i += 16 {
		hash.Reset()
		hash.Write(secret)
		hash.Write(enc[i-16 : i])
		enc = hash.Sum(enc)

		for j, b := range plaintext[i : i+16] {
			enc[i+j] ^= b
		}
	}

	return enc, nil
}
