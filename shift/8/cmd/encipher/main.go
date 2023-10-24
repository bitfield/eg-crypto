package main

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bitfield/shift"
)

func main() {
	keyHex := flag.String("key", "", "32-byte key in hexadecimal")
	flag.Parse()
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	block, err := shift.NewCipher(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	iv := make([]byte, shift.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Stdout.Write(iv)
	enc := cipher.NewCBCEncrypter(block, iv)
	plaintext = shift.Pad(plaintext, shift.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	enc.CryptBlocks(ciphertext, plaintext)
	os.Stdout.Write(ciphertext)
}
