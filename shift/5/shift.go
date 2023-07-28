package shift

import (
	"bytes"
	"errors"
)

func Encipher(plaintext []byte, key byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b + key
	}
	return ciphertext
}

func Decipher(ciphertext []byte, key byte) (plaintext []byte) {
	return Encipher(ciphertext, -key)
}

func Crack(ciphertext, crib []byte) (key byte, err error) {
	for guess := 0; guess <= 255; guess++ {
		result := Decipher(ciphertext[:len(crib)], byte(guess))
		if bytes.Equal(result, crib) {
			return byte(guess), nil
		}
	}
	return 0, errors.New("no key found")
}
