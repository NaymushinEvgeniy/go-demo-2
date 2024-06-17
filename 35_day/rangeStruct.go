package main

import (
	"html/template"
	"log"
	"os"
)

type myType struct {
	Name  string
	Count int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTmpl(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	exString := "This is obj with name {{ .Name }} and count {{ .Count }}"
	executeTmpl(exString, myType{Name: "Test", Count: 2})
}
