package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alam": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var gradesSorted []string
	for grade := range grades {
		gradesSorted = append(gradesSorted, grade)
	}
	sort.Strings(gradesSorted)
	fmt.Println(gradesSorted)
}
