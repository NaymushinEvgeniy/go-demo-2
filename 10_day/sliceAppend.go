package main

import (
	"fmt"
)

func main() {
	srcSlice := []string{"a", "b", "c"}
	s1 := append(srcSlice, "d")
	s2 := append(s1, "x")
	s3 := append(s2, "y")
	s4 := append(s3, "XX")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	s1[0] = "YY"
	fmt.Println(s4)
	fmt.Println(s1)
}
