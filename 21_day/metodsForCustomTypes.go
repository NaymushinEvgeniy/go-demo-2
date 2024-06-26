package main

import (
	"fmt"
)

type Gallons float64
type Liters float64

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel : %0.1f gallons \n", carFuel)
	fmt.Printf("Bus fuel : %0.1f litters \n", busFuel)
}
