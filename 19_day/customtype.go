package main

import (
	"fmt"
)

func main() {
	type part struct {
		count int
		descr string
	}
	type cars struct {
		model    string
		topspeed float64
	}
	var bolts part
	bolts.count = 24
	bolts.descr = "Hex bolts"
	fmt.Println(bolts)

	var volvo cars
	fmt.Println(volvo)
	volvo.model = "Volvo 360"
	volvo.topspeed = 240
	fmt.Println(volvo)

}
