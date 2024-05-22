package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	*myIntPointer = 6
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
	myIntTypePointer := reflect.TypeOf(&myInt)
	fmt.Println(myIntTypePointer)
	fmt.Println(*&myIntTypePointer)
}
