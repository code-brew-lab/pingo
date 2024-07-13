package netcore

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"math/big"
)

type ID uint16

func NewID() ID {
	id, err := rand.Int(rand.Reader, big.NewInt(65535))
	if err != nil {
		return ID(0)
	}
	return ID(id.Uint64())
}

func ParseID(bytes []byte) (ID, error) {
	if len(bytes) != 2 {
		return ID(0), errors.New("id should be 2 bytes")
	}

	be := binary.BigEndian
	id := be.Uint16(bytes[:])

	return ID(id), nil
}

func (i ID) ToUint16() uint16 {
	return uint16(i)
}
