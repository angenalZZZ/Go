package models

import (
	"encoding/hex"
	"github.com/google/uuid"
)

type ID [32]byte

func (id ID) String() string {
	return string(id[:])
}

func NewID() ID {
	dst, src := ID{}, uuid.New()
	hex.Encode(dst[:], src[:16])
	return dst
}
