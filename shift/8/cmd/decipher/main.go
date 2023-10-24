package main

import (
	"crypto/cipher"
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
	ciphertext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	iv := ciphertext[:shift.BlockSize]
	plaintext := make([]byte, len(ciphertext)-shift.BlockSize)
	dec := cipher.NewCBCDecrypter(block, iv)
	dec.CryptBlocks(plaintext, ciphertext[shift.BlockSize:])
	plaintext = shift.Unpad(plaintext, shift.BlockSize)
	os.Stdout.Write(plaintext)
}
