package main

import (
	"fmt"
)

func main() {
	numbers := [3]float64{31.2, 22.5, 33.8}
	var sum float64 = 0
	for _, value := range numbers {
		sum += value
	}
	lenArr := float64(len(numbers))
	average := sum / lenArr

	fmt.Printf("Average is %0.2f", average)
}
