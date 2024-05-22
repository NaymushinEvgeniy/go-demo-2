package main

import (
	"fmt"
)

func main() {
	titles := map[string]int{"gold": 1, "silcer": 2, "bronze": 3}
	fmt.Println(titles)
	fmt.Println(titles["gold"])
	delete(titles, "gold")
	fmt.Println(titles)
	fmt.Println(titles["gold"])
	var ranks = make(map[string]bool)
	ranks["First"] = true
	fmt.Println(ranks)
	var value bool
	var ok bool
	value, ok = ranks["First"]
	fmt.Println(value, ok)
	delete(ranks, "First")
	value, ok = ranks["First"]
	fmt.Println(value, ok)

}
