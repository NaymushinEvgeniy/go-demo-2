package main

import (
	"fmt"
)

func main() {
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[2])
	fmt.Println(boolSlice[5])
	fmt.Println(len(boolSlice))
	var intSlice []int
	var stringSlice []string

	fmt.Println(len(intSlice))
	fmt.Println(len(stringSlice))

	fmt.Printf("Int slice is %#v and string slice is %#v\n", intSlice, stringSlice)

	intSlice = append(intSlice, 27)
	fmt.Printf("Now, intslice is %#v", intSlice)
}
