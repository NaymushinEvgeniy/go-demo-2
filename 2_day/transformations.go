package main

import (
	"fmt"
)

func main() {
	var myInt int = 2
	var myFloat float64 = 3.2

	myInt = int(myFloat)

	fmt.Println(myInt)
}
