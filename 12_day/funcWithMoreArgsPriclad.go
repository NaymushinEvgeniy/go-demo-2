package main

import (
	"fmt"
)

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	mix(1, true, "tst", "a", "b")
	mix(2, false, "ssts", "c", "d")
	mix(3, true)
}
