package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bitfield/shift"
)

func main() {
	keyHex := flag.String("key", "01", "key in hexadecimal (for example 'FF')")
	flag.Parse()
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ciphertext := shift.Encipher(plaintext, key)
	os.Stdout.Write(ciphertext)
}
