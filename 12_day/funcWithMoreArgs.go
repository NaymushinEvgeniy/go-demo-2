package main

import (
	"fmt"
)

func moreArgs(argEx ...string) {
	fmt.Println(argEx)
}

func main() {
	moreArgs()
}
