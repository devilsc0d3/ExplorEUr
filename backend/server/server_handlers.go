package server

import (
	"exploreur/backend/register"
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func category(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/category.html")
	err := page.ExecuteTemplate(w, "category.html", data)
	if err != nil {
		return
	}
}

func Login(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/login.html")
	err := page.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return
	}
}

func Registration(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/registration.html")
	err := page.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		return
	}
}

func Chat(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/chat.html")
	var contents []string
	register.Db.Table("posts").Pluck("content", &contents)

	//if r.FormValue("test") != "" {
	//	print("test3")
	//	//post.AddPost(r.FormValue("post"))
	//}

	err := page.ExecuteTemplate(w, "chat.html", contents)
	if err != nil {
		return
	}
}

func Info(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	postContent := r.FormValue("postContent")
	fmt.Println("postContent", postContent)
	w.Write([]byte(postContent))
}
