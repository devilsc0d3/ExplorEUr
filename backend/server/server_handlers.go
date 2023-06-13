package server

import (
	"exploreur/backend/registerDB"
	"html/template"
	"net/http"
)

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
	} else if r.FormValue("password") != r.FormValue("confirmation") {
		http.Redirect(w, r, "/registration?error=password-not-identical", 303)
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

func RecoveringPassword(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/recovering_password.html")
	if r.FormValue("nickname") != "" && r.FormValue("password") != "" && r.FormValue("confirmation") == r.FormValue("password") {
		userError := registerDB.UpdateUserPasswordController(r.FormValue("nickname"), r.FormValue("password"))
		if userError == "" {
			http.Redirect(w, r, "/login", 303)
		}
	} else if r.FormValue("password") != r.FormValue("confirmation") {
		http.Redirect(w, r, "/registration?error=password-not-identical", 303)
	}
	err := page.ExecuteTemplate(w, "recovering_password.html", nil)
	if err != nil {
		return
	}
}
