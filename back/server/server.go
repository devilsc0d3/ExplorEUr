package main

import (
	"html/template"
	"net/http"
)

func router() {
	http.HandleFunc("/", test)
	http.HandleFunc("/category", category)
}

func test(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/home.html")
	err := page.Execute(w, nil)
	if err != nil {
		return
	}
}

func category(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/home.html")
	err := page.Execute(w, nil)
	if err != nil {
		return
	}
}
