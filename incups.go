package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Day   int
	Month int
	Year  int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	dayOne := Date{}
	err := dayOne.SetYear(2002)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dayOne.Year)
	errTwo := dayOne.SetMonth(11)
	if errTwo != nil {
		log.Fatal(errTwo)
	}
	fmt.Println(dayOne.Month)
	errTree := dayOne.SetDay(2)
	if errTree != nil {
		log.Fatal(errTree)
	}
	fmt.Println(dayOne.Month)

}
