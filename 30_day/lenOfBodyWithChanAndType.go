package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Url  string
	Size int
}

func responseSize(url string, channel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{Url: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	sizes := make(chan int)
	urls := []string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}

	for _, url := range urls {
		go responseSize(url, sizes)
	}

	for i := 0; i < len(urls); i++ {
		page <- pages
		fmt.Printf("%s: %d\n", page.Url, page.Size)
	}
}
