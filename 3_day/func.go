package main

import (
	"fmt"
)

func main() {
	repeater("test", 10)
}

func repeater(text string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(text)
	}
}
