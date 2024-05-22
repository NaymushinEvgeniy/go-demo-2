package main

import (
	"fmt"
)

func main() {
	var notes []string = make([]string, 7)
	notes[3] = "test"
	fmt.Println(notes)
	notes = make([]string, 10)
	fmt.Println(notes)

	primes := make([]int, 10)
	fmt.Println(primes)
}
