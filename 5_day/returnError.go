package main

import (
	"fmt"
	"log"
)

func paintNeeded(widht float64, height float64) (float64, error) {
	if widht <= 0 {
		return 0, fmt.Errorf(" a widht of %0.2f is invalid", widht)
	}
	if height <= 0 {
		return 0, fmt.Errorf(" a height of %0.2f is invalid", height)
	}
	area := widht * height
	area = area / 10.0
	return area, nil
}

func main() {
	amount, err := paintNeeded(10.0, -0.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Area is %0.2f", amount)
}
