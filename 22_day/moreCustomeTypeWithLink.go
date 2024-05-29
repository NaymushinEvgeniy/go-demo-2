package main

import (
	"fmt"
)

type exampleType int

func (e *exampleType) Double() {
	*e *= 2
}

func main() {
	num := exampleType(10)
	num.Double()
	fmt.Println(num)
}
