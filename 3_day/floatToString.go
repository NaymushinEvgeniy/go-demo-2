package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	buff := bufio.NewReader(os.Stdin)
	input, err := buff.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Now number is: ", number)
	fmt.Println(reflect.TypeOf(number))
}
