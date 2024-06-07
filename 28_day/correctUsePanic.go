package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Printf("You win a goat!")
	} else {
		panic("Invalid door number") // Корректная обработка потенциальных ошибок
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
