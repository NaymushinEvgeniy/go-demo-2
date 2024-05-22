package main

import (
	"fmt"
)

func main() {
	ranks := map[string]float64{
		"gold":   1,
		"silver": 2,
		"bronze": 3,
	}
	fmt.Println(ranks["gold"])

	elements := map[string]string{"H": "Hydrogenium", "Li": "Litium"}
	fmt.Println(elements["Li"])

	emptyMap := map[int]string{}
	fmt.Println(emptyMap)

	emptyMap[1] = "test1"
	fmt.Println(emptyMap)
	fmt.Println(emptyMap[1])
	fmt.Printf("In go empty map value %#v\n", emptyMap[2])

	emptyMapstr := map[string]float64{}
	emptyMapstr["test1"] = 111
	fmt.Printf("In go empty map value is %#v", emptyMapstr["test2"])
}
