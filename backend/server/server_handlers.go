package server

import (
	"exploreur/backend/register"
	"exploreur/backend/roles/user"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

type DataHub struct {
	Role        string
	Database    []Posts
	Category    []string
	IsConnected bool
}

var dataHub DataHub

var catId int

type Posts struct {
	Content  string
	Id       int
	Comments []string
	UserId int
	NicknameUser string
}

func InitRole(token string) {
	_, connectedRole, err := register.DecodeJWTToken(token)
	if err != nil {
		panic("decode token error")
	}
	switch connectedRole {
	case "":
		dataHub.Role = ""
		break
	case "user":
		dataHub.Role = "user"
		break
	case "moderator":
		dataHub.Role = "moderator"
		break
	case "administrator":
		dataHub.Role = "administrator"
		break
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")
	if dataHub.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
	}
	err := page.ExecuteTemplate(w, "home.html", dataHub)
	if err != nil {
		return
	}
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/category.html")
	if dataHub.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
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
	var categoryName []string
	register.Db.Table("categories").Pluck("name", &categoryName)
	dataHub.Category = categoryName
	err := page.ExecuteTemplate(w, "category.html", dataHub)
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
			InitRole(register.Token)
			dataHub.IsConnected = true
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
	page, err := template.ParseFiles("./front/template/chat.html")
	if err != nil {
		fmt.Println(err)
	}
	if dataHub.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
	}

	//get category_id
	compile := regexp.MustCompile(`[^/]`)
	catId, _ = strconv.Atoi(compile.FindString(r.URL.String()))

	//get info posts for put to struct
	var content []string
	var postId []int
	var message []string
	var postIdComment []int
	var userId []int

	register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("content", &content)
	register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("id", &postId)
	register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("user_id", &userId)
	register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("message", &message)
	register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("post_id", &postIdComment)

	database := ManageData(content, postId, message, postIdComment, userId)
	dataHub.Database = database
	err = page.ExecuteTemplate(w, "chat.html", dataHub)
	if err != nil {
		return
	}
}

func ManageData(content []string, postId []int, message []string, postIdComment []int, userId []int) []Posts {
	var database []Posts

	for i := 0; i < len(content); i++ {
		var temp Posts
		temp.Content = content[i]
		temp.Id = postId[i]
		temp.UserId = userId[i]

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

	for i := 0; i < len(userId); i++ {
		nickname, _ := register.GetNicknameByID(userId[i])
		for k := 0; k < len(database); k++ {
			nicknameDB, _ := register.GetNicknameByID(database[k].UserId)
			if nickname == nicknameDB {
				database[k].NicknameUser = nickname
				break
			}
		}
	}
	return database
}

// Info get info to front chat page
func Info(w http.ResponseWriter, r *http.Request) {
	if dataHub.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
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
		user.AddCommentByUserController(postID, commentContent, catId)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	register.DeleteCookie(w)
	dataHub.IsConnected = false
	http.Redirect(w, r, "/", http.StatusFound)
}

func EasterEgg(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/easter_egg.html")
	err := page.ExecuteTemplate(w, "easter_egg.html", nil)
	if err != nil {
		panic("execute template error")
	}
}

func ActivityHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/activity.html")
	if dataHub.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
	}
	err := page.ExecuteTemplate(w, "activity.html", dataHub)
	if err != nil {
		panic("execute template error")
	}
}

func RecoverHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/recovering_password.html")
	err := page.ExecuteTemplate(w, "recovering_password.html", nil)
	if err != nil {
		panic("execute template error")
	}
}
