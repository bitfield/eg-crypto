package main

import (
	"fmt"
	"io"
	"os"

	"github.com/bitfield/shift"
)

func main() {
	plaintext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ciphertext := shift.Encipher(plaintext)
	os.Stdout.Write(ciphertext)
}
