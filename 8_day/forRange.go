package main

import (
	"fmt"
)

func main() {
	var myArray [5]string = [5]string{"test1", "test2", "test3", "test4", "test5"}
	fmt.Println(myArray)
	for index, value := range myArray {
		fmt.Println("Index is ", index, "Value is ", value)
	}
}
