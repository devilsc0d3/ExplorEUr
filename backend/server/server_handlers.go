package server

import (
	"exploreur/backend/registerDB"
	"html/template"
	"net/http"
)

func Router() {
	fs := http.FileServer(http.Dir("./front/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/category", CategoryHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/registration", RegistrationHandler)
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func CategoryHandler(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/category.html")

	dataTest := []string{"place", "Tools", "information", "+"}

	err := page.ExecuteTemplate(w, "category.html", dataTest)
	if err != nil {
		return
	}
}

func LoginHandler(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/login.html")
	err := page.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return
	}
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/registration.html")
	err := page.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		return
	}
	if r.FormValue("nickname") != "" && r.FormValue("email") != "" && r.FormValue("password") != "" && r.FormValue("confirmation") == r.FormValue("password") {
		user := registerDB.AddUserController(r.FormValue("nickname"), r.FormValue("email"), r.FormValue("password"))
		if user == "" {
			http.Redirect(w, r, "/", 303)
		}
	}
}
