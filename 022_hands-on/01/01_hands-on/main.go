package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Client visit the index")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Client want a dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Richard")
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}