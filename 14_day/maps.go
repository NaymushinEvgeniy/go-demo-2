package main

import (
	"fmt"
)

func main() {
	var myMap map[string]float64
	myMap = make(map[string]float64)
	fmt.Println("My map in GO is", myMap)

	ranks := make(map[string]int)
	fmt.Println(ranks)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Litium"

	fmt.Println(elements["H"])

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
}
