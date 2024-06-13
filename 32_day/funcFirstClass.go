package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBy() {
	fmt.Println("By!")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	var myFunction func()
	myFunction = sayHi
	myFunction()

	twice(sayHi)
	twice(sayBy)
}
