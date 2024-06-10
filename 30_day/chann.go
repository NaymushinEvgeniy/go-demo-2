package main

import (
	"fmt"
	"time"
)

func greeting(myChannel chan string) {
	myChannel <- "Hi"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	time.Sleep(5 * time.Second)
	fmt.Println(<-myChannel)
}
