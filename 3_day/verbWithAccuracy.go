package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%21s | %4s\n ", "Product", "Count")
	fmt.Println("----------------------------------")
	fmt.Printf("%21s | %4d\n ", "Papper", 10)
	fmt.Printf("%20s | %4d\n ", "Apple and orange", 10)
	fmt.Printf("%20s | %4d\n ", "Other", 50)
	fmt.Printf("%20s | %4d\n ", "Hm", 500000000000)
}
