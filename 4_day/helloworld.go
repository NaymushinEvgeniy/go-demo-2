package main

import (
	"fmt"
)

func multiReturn() (int, bool, string) {
	return 228, false, "bad uk"
}

func multiReturnInt(numFloat float64, numInt int) (float64, int) {
	numFloat = numFloat / 2
	numInt = int(numFloat)
	return numFloat, numInt - 1
}

func main() {
	fmt.Println(multiReturn())
	fmt.Println(multiReturnInt(10.2, 5))
}
