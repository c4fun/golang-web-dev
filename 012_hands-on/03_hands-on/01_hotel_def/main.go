package main

import (
	"os"
	"log"
	"text/template"
)

// type hotel struct {
// 	Name string
// 	Address string
// 	City string
// 	Zip string
// 	Region string
// }

type hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name: "Malibu Beach Resort",
			Address: "Eagle Street 156",
			City: "Malibu",
			Zip: "93527",
			Region: "Southern",
		},
		hotel{
			Name: "SF Autobot",
			Address: "Cybertron 38093",
			City: "San Francisco",
			Zip: "93823",
			Region: "Central",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}