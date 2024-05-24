package main

import (
	"fmt"
)

type part struct {
	descr string
	count int
}

func showinfo(p part) {
	fmt.Println(p.descr)
	fmt.Println(p.count)
}

func minimalOrder(descr string) part {
	var p part
	p.descr = descr
	p.count = 1
	return p
}

func main() {
	var hex part
	hex.descr = "Example descr"
	hex.count = 4
	showinfo(hex)
	fmt.Println(minimalOrder("Bolt"))
	p := minimalOrder("Many Bolt")
	fmt.Println(p.count, p.descr)
}
