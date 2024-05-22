package main

import (
	"fmt"
)

func main() {
	var nilMap map[string]string
	fmt.Println(nilMap)
	nilMap["test"] = "testTT"
}
