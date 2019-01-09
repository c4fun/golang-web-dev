package main

import (
	"io"
	"log"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Client visit the index")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Client want a dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "tpl.gohtml", "xxx")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
}


func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}