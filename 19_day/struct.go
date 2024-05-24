package main

import (
	"fmt"
)

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("My first struct is %#v\n", myStruct)
	fmt.Println(myStruct)
	myStruct.number = 10
	fmt.Println(myStruct)
	myStruct.word = "test"
	fmt.Println(myStruct)
}
