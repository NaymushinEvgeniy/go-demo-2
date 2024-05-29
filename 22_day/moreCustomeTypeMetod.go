package main

import (
	"fmt"
)

type myCustomType string

func (m myCustomType) GetMoreInfo(num float64, active bool) {
	fmt.Printf("This first arg as num is: %0.2f\n", num)
	fmt.Printf("This is %v\n", active)
}

func main() {
	var exampleVar myCustomType
	exampleVar.GetMoreInfo(22.3, false)
}
