package main

import (
	"fmt"
)

type exStruct struct {
	field1 string
	field2 string
}

func main() {
	var exnum int = 2
	var linkExnum *int = &exnum
	fmt.Println(linkExnum)
	fmt.Println(*linkExnum)

	var exObj exStruct
	exObj.field1 = "test1"
	var pointer *exStruct = &exObj
	fmt.Println((*pointer).field1)
	fmt.Println(pointer.field1)
}
