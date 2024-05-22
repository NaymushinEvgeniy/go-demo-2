package main

import (
	"fmt"
	"time"
)

func main() {
	var myArray [5]time.Time
	fmt.Println(myArray)
	myArray[0] = time.Now()
	myArray[1] = time.Now()
	myArray[2] = time.Now()
	myArray[3] = time.Now()
	myArray[4] = time.Now()
	fmt.Println(myArray)
	if myArray[0] == myArray[1] {
		fmt.Println("yes, true")
	} else {
		fmt.Println("false")
	}
	fmt.Println(myArray[2].MarshalJSON())
	fmt.Println(myArray[3].Second())
}
