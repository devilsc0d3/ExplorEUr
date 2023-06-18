package server

import (
	"exploreur/backend/database/comment"
	"exploreur/backend/post"
	"exploreur/backend/register"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
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

var catId int

func Chat(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/chat.html")

	//get id
	compile := regexp.MustCompile(`[^/]`)
	catId, _ = strconv.Atoi(compile.FindString(r.URL.String()))

	//get content
	var contents []string
	register.Db.Table("posts").Where("category_id = ?", catId).Pluck("content", &contents)

	err := page.ExecuteTemplate(w, "chat.html", contents)
	if err != nil {
		return
	}
}

// Info get info to front chat page
func Info(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	//add post
	postContent := r.FormValue("postContent")
	post.AddPost(postContent, catId)

	//add comment
	if r.FormValue("comment") != "" {
		commentContent := r.FormValue("comment")
		fmt.Println(commentContent)
		comment.AddComment(commentContent)
	}

}
