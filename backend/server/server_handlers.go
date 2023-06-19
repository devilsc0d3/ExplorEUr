package server

import (
	"exploreur/backend/database/comment"
	"exploreur/backend/register"
	"exploreur/backend/roles/user"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
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
	err := page.ExecuteTemplate(w, "category.html", data)
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

var catId int

type Posts struct {
	Content []string
	Id      []int
}

type Posts2 struct {
	Content  string
	Id       int
	Comments []string
}

var text Posts

func Chat(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./front/template/chat.html")
	if err != nil {
		fmt.Println(err)
	}
	if isConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
	}

	//get category_id
	compile := regexp.MustCompile(`[^/]`)
	catId, _ = strconv.Atoi(compile.FindString(r.URL.String()))

	//get content
	var content []string
	var postId []int
	var message []string
	var postIdComment []int

	register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("content", &content)
	register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("id", &postId)
	register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("message", &message)
	register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("post_id", &postIdComment)

	database := ManageData(content, postId, message, postIdComment)
	err = page.ExecuteTemplate(w, "chat.html", database)
	if err != nil {
		return
	}
}

func ManageData(content []string, postId []int, message []string, postIdComment []int) []Posts2 {
	var database []Posts2

	for i := 0; i < len(content); i++ {
		var temp Posts2
		temp.Content = content[i]
		temp.Id = postId[i]

		database = append(database, temp)
	}

	for j := 0; j < len(message); j++ {
		for k := 0; k < len(database); k++ {
			if postIdComment[j] == database[k].Id {
				database[k].Comments = append(database[k].Comments, message[j])
				break
			}
		}
	}

	return database
}

// Info get info to front chat page
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

	//add post
	if r.FormValue("postContent") != "" {
		postErr := user.AddPostByUserController(r.FormValue("postContent"), catId)
		if postErr != "" {
			panic("post error")
		}
	}

	//add comment
	if r.FormValue("comment") != "" {
		commentContent := r.FormValue("comment")
		postID, _ := strconv.Atoi(r.FormValue("postID"))
		comment.AddComment(commentContent, postID, catId)
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
