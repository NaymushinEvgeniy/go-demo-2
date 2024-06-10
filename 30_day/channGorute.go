package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wake up!")
}

func send(myChannel chan string) {
	reportNap("sending gorouting", 2)
	fmt.Println("***sendig value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving gorouting", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
