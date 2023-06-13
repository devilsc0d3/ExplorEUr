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
	http.HandleFunc("/register", Register)
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

func Register(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("nickname") != "" && r.FormValue("email") != "" && r.FormValue("password") != "" && r.FormValue("confirmation") == r.FormValue("password") {
		userError := registerDB.AddUserController(r.FormValue("nickname"), r.FormValue("email"), r.FormValue("password"))
		if userError == "" {
			http.Redirect(w, r, "/login", 303)
		} else if userError != "" {
			if registerDB.IfNicknameExist(r.FormValue("nickname")) {
				http.Redirect(w, r, "/registration?error=nickname-already-used", 303)
			}
			if registerDB.IfEmailExist(r.FormValue("email")) {
				http.Redirect(w, r, "/registration?error=email-already-used", 303)
			}
			if !registerDB.CheckEmail(r.FormValue("email")) {
				http.Redirect(w, r, "/registration?error=email-not-valid", 303)
			}
			if !registerDB.CheckPassword(r.FormValue("password")) {
				http.Redirect(w, r, "/registration?error=password-not-valid", 303)
			}
		}
	}
	http.Redirect(w, r, "/registration", 303)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/registration.html")
	err := page.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		return
	}
}
