package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Println("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Println("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")
}
