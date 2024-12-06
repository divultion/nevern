package id

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const IdLength = 8

type Id struct {
	raw [IdLength]byte
}

func RandomId() (Id, error) {
	id_unsized, err := randomBytes(IdLength)
	if err != nil {
		return Id{raw: [IdLength]byte{}}, err
	}

	id := new([IdLength]byte)
	copy(id[:], id_unsized)

	return Id{raw: *id}, nil
}

func FromHex(input string) (Id, error) {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		return Id{}, err
	}

	if len(bytes) != IdLength {
		return Id{}, fmt.Errorf("invalid input: '%s' must decode into %d bytes but it decodes into %d", input, IdLength, len(bytes))
	}

	raw := new([IdLength]byte)
	copy(raw[:], bytes)

	return Id{raw: *raw}, nil
}

func FromRawSized(raw [IdLength]byte) Id {
	return Id{raw: raw}
}

func FromRaw(rawUnsized []byte) (Id, error) {
	if len(rawUnsized) != IdLength {
		return Id{}, fmt.Errorf("Id must be length of %d, but id received has length of %d", IdLength, len(rawUnsized))
	}

	rawId := new([IdLength]byte)
	copy(rawId[:], rawUnsized)

	id := FromRawSized(*rawId)
	return id, nil
}

func (id *Id) ToHex() string {
	return hex.EncodeToString(id.raw[:])
}

func (id *Id) ToRaw() [IdLength]byte {
	return id.raw
}

func randomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
