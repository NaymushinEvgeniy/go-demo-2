package main

import (
	"html/template"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Here is my template!\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	errTwo := tmpl.Execute(os.Stdout, nil)
	check(errTwo)
}
