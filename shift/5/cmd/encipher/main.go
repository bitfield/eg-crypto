package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/bitfield/shift"
)

func main() {
	key := flag.Int("key", 1, "shift value")
	flag.Parse()
	plaintext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ciphertext := shift.Encipher(plaintext, byte(*key))
	os.Stdout.Write(ciphertext)
}
