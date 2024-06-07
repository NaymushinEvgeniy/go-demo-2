package main

import (
	"fmt"
)

func calmDawn() {
	recover()
}

func freakOut() {
	defer calmDawn() // Откладываем вызов recover на всякий
	panic("Oh, no!") // Если в программе после этого возникает паника, отложенный вызов функции позволит провести восстановление
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
