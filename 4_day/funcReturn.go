package main

import (
	"fmt"
	"reflect"
)

func double(invalue float64) float64 {
	return invalue / 2
}

func treble(invalue float64) float64 {
	return invalue / 3
}

func main() {
	fmt.Println(double(5))
	fmt.Println(reflect.TypeOf(double(5)))
	result := double(7)
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
	fmt.Println(treble(10))
}
