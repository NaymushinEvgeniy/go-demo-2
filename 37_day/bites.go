package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("false && false == %t\n", false && false)
	fmt.Printf("false && true == %t\n", false && true)
	fmt.Printf("true && true == %t\n", true && true)

	fmt.Printf("%b & %b == %b\n", 0, 0, 0&0)
	fmt.Printf("%b & %b == %b\n", 0, 1, 0&1)
	fmt.Printf("%b & %b == %b\n", 1, 1, 1&1)

	fmt.Println(170 & 15)
	fmt.Println(10 & 7)
	fmt.Println(100 & 45)

	fmt.Printf("%02b\n", 1)
	fmt.Printf("%02b\n", 2)
	fmt.Printf("%02b\n", 3)
	fmt.Printf("%02b\n", 1&3)
	fmt.Printf("%02b\n", 2&3)

	fmt.Printf("%08b\n", 170)
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", 170&15)
	fmt.Printf("%08b\n", 100&45)

	fmt.Println(os.O_RDONLY, os.O_WRONLY, os.O_RDWR, os.O_CREATE, os.O_APPEND)
	fmt.Printf("%016b\n", os.O_RDONLY)
	fmt.Printf("%016b\n", os.O_WRONLY)
	fmt.Printf("%016b\n", os.O_RDWR)
	fmt.Printf("%016b\n", os.O_CREATE)
	fmt.Printf("%016b\n", os.O_APPEND)

	fmt.Printf("%016b\n", os.O_WRONLY|os.O_CREATE)
	fmt.Printf("%016b\n", os.O_WRONLY|os.O_CREATE|os.O_APPEND)
}
