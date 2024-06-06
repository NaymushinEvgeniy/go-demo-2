package main

import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("It's defer for one func")
	two()
}

func two() {
	defer fmt.Println("It's defer for two func")
	tree()
}

func tree() {
	panic("This is call stack for me")
}
