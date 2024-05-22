package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(reflect.TypeOf(now))
	fmt.Println(now)
	var year int = now.Year()

	fmt.Println(year)
}
