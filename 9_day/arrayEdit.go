package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func arrEdit(filename string) ([3]float64, error) {
	var numbers [3]float64
	i := 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

func main() {
	rezult, err := arrEdit("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rezult)
}
