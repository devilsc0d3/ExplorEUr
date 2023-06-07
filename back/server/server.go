package main

import (
	"html/template"
	"net/http"
)

func router() {
	http.HandleFunc("/", home)
	http.HandleFunc("/category", category)
}

func home(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../../front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func category(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../../front/template/category.html")

	dataTest := []string{"place", "Tools", "information", "+"}

	err := page.ExecuteTemplate(w, "category.html", dataTest)
	if err != nil {
		return
	}
}
