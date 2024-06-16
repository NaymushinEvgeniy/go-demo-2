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
	myTemplate := "Start of template\n Value of template: {{.}}\n End of template\n"
	tmpl, err := template.New("test").Parse(myTemplate)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, false)
	check(err)
	err = tmpl.Execute(os.Stdout, 344334)
}
