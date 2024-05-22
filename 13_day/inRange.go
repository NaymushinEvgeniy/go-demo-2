package main

import "fmt"

func inRange(start float64, stop float64, numbers ...float64) {
	var resultSlice []float64
	for _, num := range numbers {
		if num >= start && num <= stop {
			resultSlice = append(resultSlice, num)
		}
	}
	fmt.Println(resultSlice)
}

func main() {
	inRange(10, 20, 11, -1, -200, 20, 22, 12, 65)
}
