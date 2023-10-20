package shift

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"fmt"
)

const BlockSize = 32

var ErrKeySize = errors.New("shift: invalid key size")

type shiftCipher struct {
	key [BlockSize]byte
}

func NewCipher(key []byte) (cipher.Block, error) {
	if len(key) != BlockSize {
		return nil, fmt.Errorf("%w %d (must be %d)", ErrKeySize,
			len(key), BlockSize)
	}
	return &shiftCipher{
		key: [BlockSize]byte(key),
	}, nil
}

func (c *shiftCipher) Encrypt(dst, src []byte) {
	for i, b := range src {
		dst[i] = b + c.key[i]
	}
}

func (c *shiftCipher) Decrypt(dst, src []byte) {
	for i, b := range src {
		dst[i] = b - c.key[i]
	}
}

func (c shiftCipher) BlockSize() int {
	return BlockSize
}

func Pad(data []byte, blockSize int) []byte {
	n := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(n)}, n)
	return append(data, padding...)
}

func Unpad(data []byte, blockSize int) []byte {
	n := int(data[len(data)-1])
	return data[:len(data)-n]
}
