package main

import (
	"html/template"
	"net/http"
	"path"
)

type Element struct {
	Name   string
	Number int
}

func main() {
	http.HandleFunc("/", displayName)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}
}

func displayName(w http.ResponseWriter, r *http.Request) {
	element := Element{"Unique", 1}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, element); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
