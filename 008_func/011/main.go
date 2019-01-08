package main

import (
	"os"
	"text/template"
)


func main() {
	tpl := template.Must(template.New("someth").Parse(`here is my template, yo`))
	tpl.ExecuteTemplate(os.Stdout, "Someth", nil)

}
