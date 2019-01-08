package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type resturant struct {
	Name, Address, City string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	resturants := []resturant {
		resturant{
			Name: "Jianqiao Diner",
			Address: "Chongqing Internet Industrial Park 101",
			City: "Chongqing",
			Menu: menu {
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:	"Xiao Mian",
							Descrip:"Spicy",
							Price:  6.5,
						},
						item{
							Name:	"Dou Jiang",
							Descrip:"Smooth",
							Price:  2.2,
						},
					},
				},
			},
		},
		resturant{
			Name: "Malibu Cuisine",
			Address: "Eagle Street 38203",
			City: "Malibu",
			Menu: menu {
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{
							Name:    "Oatmeal",
							Descrip: "yum yum",
							Price:   4.95,
						},
						item{
							Name:    "Cheerios",
							Descrip: "American eating food traditional now",
							Price:   3.95,
						},
						item{
							Name:    "Juice Orange",
							Descrip: "Delicious drinking in throat squeezed fresh",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{
							Name:    "Hamburger",
							Descrip: "Delicous good eating for you",
							Price:   6.95,
						},
						item{
							Name:    "Cheese Melted Sandwich",
							Descrip: "Make cheese bread melt grease hot",
							Price:   3.95,
						},
						item{
							Name:    "French Fries",
							Descrip: "French eat potatoe fingers",
							Price:   2.95,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{
							Name:    "Pasta Bolognese",
							Descrip: "From Italy delicious eating",
							Price:   7.95,
						},
						item{
							Name:    "Steak",
							Descrip: "Dead cow grilled bloody",
							Price:   13.95,
						},
						item{
							Name:    "Bistro Potatoe",
							Descrip: "Bistro bar wood American bacon",
							Price:   6.95,
						},
					},
				},
			},
		},
	}


	err := tpl.Execute(os.Stdout, resturants)
	if err != nil {
		log.Fatalln(err)
	}
}
