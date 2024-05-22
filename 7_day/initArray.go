package main

import (
	"fmt"
)

func main() {
	var myNewArr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myNewArr[2])

	myNewArrTwo := [5]int{10, 20, 30, 40}
	fmt.Println(myNewArrTwo[2])
	fmt.Println(myNewArrTwo[4])

	fmt.Printf("In go is %#v \n", myNewArrTwo)

	index := 0
	fmt.Println(myNewArr[index])
	index = 2
	fmt.Println(myNewArr[index])
	fmt.Println(len(myNewArrTwo))
}
