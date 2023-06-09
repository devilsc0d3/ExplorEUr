package server

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func category(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./../front/template/category.html")

	dataTest := []string{"plAce", "Tools", "informatiOn", "+"}

	err := page.ExecuteTemplate(w, "category.html", dataTest)
	if err != nil {
		return
	}
}
