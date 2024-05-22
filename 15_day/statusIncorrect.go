package main

import (
	"fmt"
)

func status(name string) {
	results := map[string]float64{"Joe": 78.5, "Lili": 30.0}
	result := results[name]
	if result < 60 {
		fmt.Printf("%s is failed\n", name)
	}
}
func main() {
	status("Lili")
	status("Carl")
}
