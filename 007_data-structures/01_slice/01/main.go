package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	// sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	got := []string{"Danerys", "Little Finger", "Arya Stark", "Jamie Lannister"}

	err := tpl.Execute(os.Stdout, got)
	if err != nil {
		log.Fatalln(err)
	}
}
