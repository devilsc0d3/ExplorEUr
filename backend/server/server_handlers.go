package server

import (
	"exploreur/backend/database/category"
	"exploreur/backend/register"
	"exploreur/backend/roles/moderator"
	"exploreur/backend/roles/user"
	"exploreur/backend/structure"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

//type DataHub struct {
//	Role                 string
//	Database             []Posts
//	Category             []string
//	IsConnected          bool
//	ReportPostContent    string
//	ReportCommentContent string
//}

//var dataHub DataHub

var catId int

//type Posts struct {
//	Content      string
//	Id           int
//	Comments     []Comment
//	UserId       int
//	NicknamePost string
//	CountLike    int
//}
//
//type Comment struct {
//	Message         string
//	PostId          int
//	NicknameComment string
//}

func InitRole(token string) {
	_, connectedRole, err := register.DecodeJWTToken(token)
	if err != nil {
		panic("decode token error")
	}
	switch connectedRole {
	case "":
		structure.DataHub1.Role = ""
		break
	case "user":
		structure.DataHub1.Role = "user"
		break
	case "moderator":
		structure.DataHub1.Role = "moderator"
		break
	case "admin":
		structure.DataHub1.Role = "admin"
		break
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	}

	if structure.DataHub1.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
	}
	err := page.ExecuteTemplate(w, "home.html", structure.DataHub1)
	if err != nil {
		return
	}
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("./front/template/category.html")
	if err != nil {
		panic("failed to parse template")
	}
	if structure.DataHub1.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie retrieval error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
	}
	var categoryName []string
	register.Db.Table("categories").Pluck("name", &categoryName)
	structure.DataHub1.Category = categoryName
	fmt.Println(structure.DataHub1)
	err = page.ExecuteTemplate(w, "category.html", structure.DataHub1)
	if err != nil {
		panic("failed to execute template")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/login.html")
	if r.FormValue("nickname") != "" && r.FormValue("password") != "" {
		isok, usr := register.CheckNicknameAndPassword(r.FormValue("nickname"), r.FormValue("password"))
		if isok {
			var token string
			var err error
			if r.FormValue("remember-me") == "1" {
				token, err = register.CreateJWTTokenRememberMe(usr.Nickname, usr.Role)
			} else {
				token, err = register.CreateJWTToken(usr.Nickname, usr.Role)
			}
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			register.CreateCookie(w, token)
			register.Token = token
			InitRole(register.Token)
			structure.DataHub1.IsConnected = true
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
	if structure.DataHub1.IsConnected {
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
	var userIdComment []int
	var postIdLike []int

	register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("content", &content)
	register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("id", &postId)
	register.Db.Table("posts").Where("category_id = ?", catId).Pluck("user_id", &userId)
	register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("message", &message)
	register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("post_id", &postIdComment)
	register.Db.Table("comments").Where("category_id = ?", catId).Pluck("user_id", &userIdComment)
	register.Db.Table("like_posts").Where("is_like = ?", true).Pluck("post_id", &postIdLike)

	// order select
	if r.FormValue("order") == "desc" {

		register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("content", &content)
		register.Db.Table("posts").Where("category_id = ?", catId).Order("created_at DESC").Pluck("id", &postId)
		register.Db.Table("posts").Where("category_id = ?", catId).Pluck("user_id", &userId)
		register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("message", &message)
		register.Db.Table("comments").Where("category_id = ?", catId).Order("created_at DESC").Pluck("post_id", &postIdComment)
		register.Db.Table("comments").Where("category_id = ?", catId).Pluck("user_id", &userIdComment)
		register.Db.Table("like_posts").Where("is_like = ?", true).Pluck("post_id", &postIdLike)

	} else if r.FormValue("order") == "asc" {

		register.Db.Table("posts").Where("category_id = ?", catId).Pluck("content", &content)
		register.Db.Table("posts").Where("category_id = ?", catId).Pluck("id", &postId)
		register.Db.Table("posts").Where("category_id = ?", catId).Pluck("user_id", &userId)
		register.Db.Table("comments").Where("category_id = ?", catId).Pluck("message", &message)
		register.Db.Table("comments").Where("category_id = ?", catId).Pluck("post_id", &postIdComment)
		register.Db.Table("comments").Where("category_id = ?", catId).Pluck("user_id", &userIdComment)
		register.Db.Table("like_posts").Where("is_like = ?", true).Pluck("post_id", &postIdLike)
	}

	database := ManageData(content, postId, message, postIdComment, userId, userIdComment, postIdLike)
	structure.DataHub1.Database = database

	err = page.ExecuteTemplate(w, "chat.html", structure.DataHub1)
	if err != nil {
		return
	}
}

func ManageData(content []string, postId []int, message []string, postIdComment []int, userId []int, userIdComment []int, postIdLike []int) []structure.Posts {
	var database []structure.Posts
	countLike := CountLike(postIdLike)
	for i := 0; i < len(content); i++ {
		var temp structure.Posts
		temp.Content = content[i]
		temp.Id = postId[i]
		temp.UserId = userId[i]
		temp.NicknamePost, _ = register.GetNicknameByID(userId[i])

		for idPost, nbrLike := range countLike {
			if idPost == postId[i] {
				temp.CountLike = nbrLike
			}
		}
		database = append(database, temp)
	}

	for i := 0; i < len(message); i++ {
		var temp2 structure.Comment
		for j := 0; j < len(database); j++ {
			if postIdComment[i] == database[j].Id {
				temp2.Message = message[i]
				temp2.PostId = postIdComment[i]
				temp2.NicknameComment, _ = register.GetNicknameByID(userIdComment[i])
				database[j].Comments = append(database[j].Comments, temp2)
				break
			}
		}
	}

	return database
}

func CountLike(postId []int) map[int]int {
	counts := make(map[int]int)

	for _, idPost := range postId {
		counts[idPost]++
	}

	return counts
}

// Info get info to front chat page
func Info(w http.ResponseWriter, r *http.Request) {
	if structure.DataHub1.IsConnected {
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

	//add category
	if r.FormValue("categoryName") != "" {
		category.AddCategory(r.FormValue("categoryName"))
	}

	//init textReport
	if structure.DataHub1.ReportPostContent == "" {
		structure.DataHub1.ReportPostContent = r.FormValue("textPostReport")
	}

	//init commentReport
	if structure.DataHub1.ReportCommentContent == "" {
		structure.DataHub1.ReportCommentContent = r.FormValue("textCommentReport")
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

	//add like/dislike post
	if r.FormValue("like") != "" {
		like := r.FormValue("like")
		dislike := r.FormValue("dislike")
		postID, _ := strconv.Atoi(r.FormValue("postID"))
		if like == "true" {
			user.AddLikePostByUserController(postID)
		}
		if dislike == "true" {
			user.AddDislikePostByUserController(postID)
		}

	}

	//report post
	if r.FormValue("reportPostButton") != "" {
		postID, _ := strconv.Atoi(r.FormValue("reportPostButton"))
		var userId int
		register.Db.Table("posts").Where("category_id = ?", catId).Pluck("user_id", &userId)
		nicknameUser, _ := register.GetNicknameByID(userId)
		moderator.ReportPostByModeratorController(postID, nicknameUser, structure.DataHub1.ReportPostContent, catId)
	}

	//report comment
	if r.FormValue("reportCommentButton") != "" {
		postID, _ := strconv.Atoi(r.FormValue("reportPostButton"))
		var userId int
		register.Db.Table("posts").Where("category_id = ?", catId).Pluck("user_id", &userId)
		nicknameUser, _ := register.GetNicknameByID(userId)
		moderator.ReportPostByModeratorController(postID, nicknameUser, structure.DataHub1.ReportPostContent, catId)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	register.DeleteCookie(w)
	structure.DataHub1.IsConnected = false
	http.Redirect(w, r, "/", http.StatusFound)
}

func EasterEgg(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/easter_egg.html")
	err := page.ExecuteTemplate(w, "easter_egg.html", nil)
	if err != nil {
		panic("execute template error")
	}
}

func Error(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/error.html")
	err := page.ExecuteTemplate(w, "error.html", nil)
	if err != nil {
		panic("execute template error")
	}
}

func ActivityHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/activity.html")
	if structure.DataHub1.IsConnected {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic("cookie recuperation error")
		}
		register.Token = cookie.Value
		InitRole(register.Token)
	}
	err := page.ExecuteTemplate(w, "activity.html", structure.DataHub1)
	if err != nil {
		panic("execute template error")
	}
}

func RecoverHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/recovering_password.html")
	if r.FormValue("nickname") != "" && r.FormValue("password") != "" && r.FormValue("confirmation") == r.FormValue("password") {
		register.UpdateUserPasswordController(r.FormValue("nickname"), r.FormValue("password"))
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	err := page.ExecuteTemplate(w, "recovering_password.html", nil)
	if err != nil {
		panic("execute template error")
	}
}
