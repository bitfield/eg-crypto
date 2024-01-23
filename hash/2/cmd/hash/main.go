package main

import (
	"fmt"
	"io"
	"os"

	"github.com/bitfield/hash"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%x\n", hash.SumHash(data))
}
