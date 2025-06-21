package utils

import (
	"crypto/rand"
	"fmt"

	"github.com/mr-tron/base58/base58"
)

const (
	VERSION   = byte(1)
	MAX_BYTE  = 255
	MIN_BYTE  = 10
	SEPERATOR = "_"
)

type Key struct {
	VERSION    byte
	Prefix     *string
	RandomByte []byte
}

func NewKey(prefix *string, length int) (*Key, error) {
	if length > MAX_BYTE || length < MIN_BYTE {
		return nil, fmt.Errorf("invald byte length")
	}

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return &Key{
		VERSION:    VERSION,
		Prefix:     prefix,
		RandomByte: b,
	}, nil
}

func (k *Key) ToString() string {
	buf := make([]byte, 2+len(k.RandomByte))
	buf[0] = k.VERSION
	buf[1] = byte(len(k.RandomByte))

	copy(buf[2:], k.RandomByte)

	encoded := base58.Encode(buf)

	if k.Prefix != nil {
		return *k.Prefix + SEPERATOR + encoded
	}
	return encoded
}
