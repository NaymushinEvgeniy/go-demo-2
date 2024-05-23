package main

import (
	"fmt"
)

func main() {
	grades := map[string]float64{"Genry": 22.6, "Charlie": 45.3, "Joe": 10.3}
	for name := range grades {
		fmt.Println(name)
	}
	for _, grade := range grades {
		fmt.Println(grade)
	}
}
