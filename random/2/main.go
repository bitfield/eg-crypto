package main

import (
	"fmt"
	"time"
)

var (
	seed1 = time.Now().Second()
	seed2 = time.Now().Second()
)

func main() {
	fmt.Println("Seeds:", seed1, seed2)
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
