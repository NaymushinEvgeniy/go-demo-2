package main

import (
	"fmt"
)

type Date struct {
	Day   int
	Month int
	Year  int
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	dayOne := Date{}
	dayOne.SetYear(2002)
	dayOne.SetMonth(11)
	dayOne.SetDay(2)
	dayTwo := Date{}
	dayTwo.SetYear(1992)
	dayTwo.SetMonth(2)
	dayTwo.SetDay(11)
	fmt.Println(dayOne)
	fmt.Println(dayTwo)
}
