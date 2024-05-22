package main

import (
	"fmt"
)

func main() {
	srcArray := [5]string{"a", "b", "c", "d"}

	sliceOne := srcArray[:3]
	sliceTwo := srcArray[2:]
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)
	srcArray[2] = "x"
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)

	staticSlice := make([]int, 10)
	fmt.Println(staticSlice)
}
