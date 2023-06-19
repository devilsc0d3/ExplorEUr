package user

import (
	"exploreur/backend/post"
	"exploreur/backend/register"
)

func AddPostByUserController(input string, categoryID int) string {
	if !CheckInsults(input) {
		return "there is at least one insult in the text"
	}
	if !CheckLength(input) {
		return "the text are too long"
	}
	nickname, _, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	id, _ := register.GetIDByNickname(nickname)
	post.AddPost(input, id, categoryID)
	return ""
}

func DeletePostByUserController() {
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "user" {
		id, _ := register.GetIDByNickname(nickname)
		post.DeletePost(id)
	}
}
