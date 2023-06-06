package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func router() {
	http.HandleFunc("/", test)
	http.HandleFunc("/category", category)
}

func test(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")
	err := page.Execute(w, nil)
	if err != nil {
		return
	}
}

func category(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")
	err := page.Execute(w, nil)
	if err != nil {
		return
	}
}

const port = "8080"

func Server() {
	router()
	fmt.Println("Listening on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return
	}
}
