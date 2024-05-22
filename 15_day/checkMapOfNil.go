package main

import (
	"fmt"
)

func main() {
	counters := map[string]int{"a": 10, "b": 30}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["x"]
	fmt.Println(value, ok)
}
