package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffePot := CoffeePot("LuxBrew")
	fmt.Println(coffePot)
	fmt.Print(coffePot, "\n")
	fmt.Printf("%s", coffePot)
}
