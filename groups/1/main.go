package main

import "fmt"

func main() {
	input := []byte("This message takes exactly 35 bytes")
	groups(input, 5)
}

func groups(data []byte, size int) {
	for len(data) > 0 {
		fmt.Printf("%s\n", data[:size])
		data = data[size:]
	}
}
