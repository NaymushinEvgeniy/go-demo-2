package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(math.Floor(22.5))
	fmt.Println(strings.Title("head string"))
	fmt.Println(reflect.TypeOf(22))
	fmt.Println(reflect.TypeOf("str"))
	fmt.Println(reflect.TypeOf('Х'))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf('>'))
	fmt.Println(reflect.TypeOf(5.2))
}
