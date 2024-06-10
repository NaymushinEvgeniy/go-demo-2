package main

import "fmt"

func abc(myChannel1 chan string) {
	myChannel1 <- "a"
	myChannel1 <- "b"
	myChannel1 <- "c"
}

func def(myChannel2 chan string) {
	myChannel2 <- "d"
	myChannel2 <- "e"
	myChannel2 <- "f"
}

func main() {
	myChannel1 := make(chan string)
	myChannel2 := make(chan string)
	go abc(myChannel1)
	go def(myChannel2)
	fmt.Print(<-myChannel1)
	fmt.Print(<-myChannel2)
	fmt.Print(<-myChannel1)
	fmt.Print(<-myChannel2)
	fmt.Print(<-myChannel1)
	fmt.Print(<-myChannel2)
}
