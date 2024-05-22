// Программа по угадыванию чисел

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Введите число ")
	buff := bufio.NewReader(os.Stdin)
	input, err := buff.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ваше число: " + input)
}
