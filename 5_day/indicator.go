package main

import (
	"fmt"
	"reflect"
)

func main() {
	var test1 int
	fmt.Println(&test1)
	fmt.Println(reflect.TypeOf(&test1))
	var test2 string
	fmt.Println(&test2)
	fmt.Println(reflect.TypeOf(&test2))
}
