package main

import (
	"fmt"
)

type subscriber struct {
	name     string
	rate     float64
	isActive bool
}

func getinfo(s *subscriber) {
	fmt.Println("Name of sub is:", s.name)
	fmt.Println(s.name, "has rate is:", s.rate)
	fmt.Println(s.name, "is active? :", s.isActive)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 1.0
	s.isActive = false
	return &s
}

func main() {
	var sub1 = defaultSubscriber("Ivan")
	getinfo(sub1)
	var sub2 subscriber
	sub2.name = "Petr"
	sub2.rate = 5.0
	sub2.isActive = true
	getinfo(&sub2)
}
