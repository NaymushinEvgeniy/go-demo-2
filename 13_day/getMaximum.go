package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) {
	specVal := math.Inf(-1)
	for _, num := range numbers {
		if num > specVal {
			specVal = num
		}
	}
	fmt.Println(specVal)
}

func main() {
	maximum(20.2, 30.2, 10.3, 4.5)
}
