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
	page, _ := template.ParseFiles("./../../front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func category(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../../front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}
