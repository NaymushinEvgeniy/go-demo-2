package main

import (
	"fmt"
)

func main() {
	mySrcArr := []int{1, 2, 5, 8}
	mySlice := mySrcArr[1:3]
	fmt.Println(mySlice)
	mySliceSeconde := mySrcArr[2:]
	fmt.Println(mySliceSeconde)
}
