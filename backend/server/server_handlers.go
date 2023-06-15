package server

import (
	"exploreur/backend/register"
	_ "exploreur/backend/register"
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/login.html")
	if r.FormValue("nickname") != "" && r.FormValue("password") != "" {
		isok, user := register.CheckNicknameAndPassword(r.FormValue("nickname"), r.FormValue("password"))
		if isok {
			token, err := register.CreateJWTToken(user.Nickname, user.Role)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			register.CreateCookie(w, token)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
	err := page.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("nickname") != "" && r.FormValue("email") != "" && r.FormValue("password") != "" && r.FormValue("confirmation") == r.FormValue("password") {
		userError := register.AddUserController(r.FormValue("nickname"), r.FormValue("email"), r.FormValue("password"))
		if userError == "" {
			http.Redirect(w, r, "/login", 303)
		} else if userError != "" {
			if register.IfNicknameExist(r.FormValue("nickname")) {
				http.Redirect(w, r, "/registration?error=nickname-already-used", 303)
			}
			if register.IfEmailExist(r.FormValue("email")) {
				http.Redirect(w, r, "/registration?error=email-already-used", 303)
			}
			if !register.CheckEmail(r.FormValue("email")) {
				http.Redirect(w, r, "/registration?error=email-not-valid", 303)
			}
			if !register.CheckPassword(r.FormValue("password")) {
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

func Chat(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/chat.html")
	err := page.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	register.DeleteCookie(w)
	http.Redirect(w, r, "/", http.StatusFound)
}
