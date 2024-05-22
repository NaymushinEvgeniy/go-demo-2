package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Test string is %s\n", "test")
	fmt.Printf("Test number is %d\n", 10)
	fmt.Printf("Test string and number is %s and %d\n", "test", 10)
	fmt.Printf("Test bool is %t\n", true)
	fmt.Printf("Test sighn is %v %v %v %v\n", 10, "\t", "test", false)
	fmt.Printf("Test sighn is %v %#v %v %v\n", 10, "\t", "test", false)
	fmt.Printf("Test type is %T\n", false)
}
