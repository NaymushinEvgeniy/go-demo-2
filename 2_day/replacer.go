package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var text string = "G# r#ck!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(text)
	fmt.Println(fixed)
	fmt.Println(reflect.TypeOf(replacer))
}
