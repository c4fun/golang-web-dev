package main

import (
	"os"
	"log"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := region{
		Region: "Central",
		Hotels: []hotel{
			hotel{
				Name: "Malibu Beach Resort",
				Address: "Eagle Street 156",
				City: "Malibu",
				Zip: "93527",
			},
			hotel{
				Name: "SF Autobot",
				Address: "Cybertron 38093",
				City: "San Francisco",
				Zip: "93823",
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}