package main

import (
	"log"
	"os"
	"text/template"
	"time"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func yearMonthDay(t time.Time) string {
	return t.Format("2006年01月02日")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"fdateYMD": yearMonthDay,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	t := time.Now()
	fmt.Println(t)
	tf := t.Format(time.RFC3339)
	fmt.Println(tf)
}
