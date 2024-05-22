package main

import (
	"fmt"
)

func status(name string) {
	results := map[string]float64{"Joe": 78.5, "Lili": 30.0}
	var ok bool
	result, ok := results[name]
	if !ok {
		fmt.Printf("%s - This entry not in map\n", name)
	} else {
		if result < 60 {
			fmt.Printf("%s is failed exam!\n", name)
		}
	}
}
func main() {
	status("Lili")
	status("Carl")
}
