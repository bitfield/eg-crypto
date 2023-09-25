package main

import (
	"fmt"
	"time"
)

var (
	seed1 = 0
	seed2 = 1
)

func main() {
	for range time.Tick(time.Second) {
		number := seq()
		fmt.Print(number, ", ")
	}
}

func seq() int {
	tmp := seed1 + seed2
	seed1 = seed2
	seed2 = tmp
	return tmp % 10
}
