package main

import (
	"fmt"
)

func main() {
	var myNewVar string
	var myNewVarSecond, myNewVarThird int

	fmt.Println(myNewVar)
	fmt.Println(myNewVarSecond)
	fmt.Println(myNewVarThird)
	fmt.Println("-----\n")

	myNewVar = "test"
	fmt.Println(myNewVar)

	myNewVarSecond = 'А'
	fmt.Println(myNewVarSecond)

	var myNewVarFourth, myNewVarFifth float32

	myNewVarFifth, myNewVarFourth = 5.228, 4.228

	fmt.Println(myNewVarFourth)
	fmt.Println(myNewVarFifth)

	var myNewVarSixth = "sixth test"
	fmt.Println(myNewVarSixth)

	var myNewVarSeventh float32
	println(myNewVarSeventh)
}
