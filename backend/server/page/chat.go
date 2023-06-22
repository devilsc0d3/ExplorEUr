package page

import (
	"exploreur/backend/register"
	"exploreur/backend/structure"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

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
		structure.InitRole(register.Token)
	}

	//get category_id
	compile := regexp.MustCompile(`[^/]`)
	catId, _ := strconv.Atoi(compile.FindString(r.URL.String()))

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

	database := structure.ManageData(content, postId, message, postIdComment, userId, userIdComment, postIdLike)
	structure.DataHub1.Database = database

	err = page.ExecuteTemplate(w, "chat.html", structure.DataHub1)
	if err != nil {
		return
	}
}
