package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := "My string"
	fmt.Println(test)
	fmt.Println(reflect.TypeOf(test))

	testTwo := 2560000000000000000
	fmt.Println(testTwo)
	fmt.Println(reflect.TypeOf(testTwo))

	testFirth := 0.0
	fmt.Println(testFirth)
	fmt.Println(reflect.TypeOf(testFirth))

	testFourth := false
	fmt.Println(reflect.TypeOf(testFourth))
}
