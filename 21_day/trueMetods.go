package main

import (
	"fmt"
)

type MyType string

func (m MyType) SayHello() {
	fmt.Printf("Hi from %s !\n", m)
}

func main() {
	value := MyType("Test")
	value.SayHello()
	otherValue := MyType("Other test")
	otherValue.SayHello()
}
