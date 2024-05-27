package main

import (
	"fmt"
	"reflect"
)

type Liters float64

type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)

	var carFuelTwo Gallons
	var busFuelTwo Liters
	carFuelTwo = 10.0
	busFuelTwo = 240.0
	fmt.Println(carFuelTwo, busFuelTwo)

	fmt.Println(reflect.TypeOf(carFuel))
	fmt.Println(reflect.TypeOf(carFuelTwo))

	var carFuelTree Gallons
	var busFuelTree Liters
	carFuelTree = Gallons(Liters(40.0) * 0.264)
	busFuelTree = Liters(Gallons(63.0) * 3.785)
	fmt.Println(carFuelTree)
	fmt.Println(reflect.TypeOf(carFuelTree))
	fmt.Println(busFuelTree)
	fmt.Println(reflect.TypeOf(busFuelTree))

	fmt.Println(Gallons(10.2) + Gallons(22.2))
	fmt.Println(Liters(20.2) * Liters(2.2))
	fmt.Println(Liters(1.2) == 1.2)
}
