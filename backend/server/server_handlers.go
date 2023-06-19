package server

import (
	"exploreur/backend/register"
	"exploreur/backend/roles/user"
	"html/template"
	"net/http"
)

var isConnected = false

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")
	if isConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
	}
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/category.html")
	if isConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		_, role, err := register.DecodeJWTToken(cookie.Value)
		if err != nil {
			panic("decode token error")
		}
		if role == "admin" {
			// il faut que ce soit des bouttons en js non ?
			dataTest := []string{"place", "Tools", "information", "+"}
			err := page.ExecuteTemplate(w, "category.html", dataTest)
			if err != nil {
				panic("execute template error")
			}
		}
	}
	dataTest := []string{"place", "Tools", "information"}
	err := page.ExecuteTemplate(w, "category.html", dataTest)
	if err != nil {
		panic("execute template error")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/login.html")
	if r.FormValue("nickname") != "" && r.FormValue("password") != "" {
		isok, user := register.CheckNicknameAndPassword(r.FormValue("nickname"), r.FormValue("password"))
		if isok {
			var token string
			var err error
			if r.FormValue("remember-me") == "1" {
				token, err = register.CreateJWTTokenRememberMe(user.Nickname, user.Role)
			} else {
				token, err = register.CreateJWTToken(user.Nickname, user.Role)
			}
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			register.CreateCookie(w, token)
			register.Token = token
			isConnected = true
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
		if userError != "" {
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
		} else if userError == "" {
			http.Redirect(w, r, "/login", 303)
		}
	}
	http.Redirect(w, r, "/registration", 303)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/registration.html")
	err := page.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		panic("execute template error")
	}
}

func Chat(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/chat.html")
	if isConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
	}
	var contents []string
	register.Db.Table("posts").Pluck("content", &contents)
	err := page.ExecuteTemplate(w, "chat.html", contents)
	if err != nil {
		return
	}
}

func Info(w http.ResponseWriter, r *http.Request) {
	if isConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
	}
	err := r.ParseForm()
	if err != nil {
		return
	}

	if r.FormValue("postContent") != "" {
		postErr := user.AddPostByUserController(r.FormValue("postContent"))
		if postErr != "" {
			panic("post error")
		}
	}

	if r.FormValue("postID") != "" {
		postErr := user.AddPostByUserController(r.FormValue("postID"))
		if postErr != "" {
			panic("post error")
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	register.DeleteCookie(w)
	isConnected = false
	http.Redirect(w, r, "/", http.StatusFound)
}

func EasterEgg(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/easter_egg.html")
	err := page.ExecuteTemplate(w, "easter_egg.html", nil)
	if err != nil {
		return
	}
}
