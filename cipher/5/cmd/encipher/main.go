package main

import (
	"cipher"
	
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	fmt.Println(cipher.Encipher(input))
}
