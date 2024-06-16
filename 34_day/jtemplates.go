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

func executeTmpl(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	myTemplate := "Start of template\n Value of template: {{.}}\n End of template\n"
	executeTmpl(myTemplate, "ABV")
	executeTmpl("This is dot {{ if . }}Yes is true!{{end}}\n", true)
	executeTmpl("This is dot {{ if . }}Yes is true!{{end}}\n", false)
	exLoop := "This is loop {{ . }}\n in loop:\n {{ range .}} value: {{.}}\n {{ end}} After loop :{{ . }}\n"
	executeTmpl(exLoop, []string{"do", "re", "mi"})
}
