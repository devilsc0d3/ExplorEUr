package main

import (
	"html/template"
	"net/http"
)

func router() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/category", CategoryHandler)
	http.HandleFunc("/login", LoginHandler)
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../../front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func CategoryHandler(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../../front/template/category.html")

	dataTest := []string{"place", "Tools", "information", "+"}

	err := page.ExecuteTemplate(w, "category.html", dataTest)
	if err != nil {
		return
	}
}

func LoginHandler(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../../front/template/login.html")
	err := page.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return
	}
}
