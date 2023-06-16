package user

import (
	"exploreur/backend/post"
	"exploreur/backend/register"
)

func AddPostByUserController(input string) string {
	if !CheckInsults(input) {
		return "there is at least one insult in the text"
	}
	if !CheckLength(input) {
		return "the text are too long"
	}
	nickname, _, _ := register.DecodeJWTToken(register.Token)
	id, _ := register.GetIDByNickname(nickname)
	post.AddPost(id, input)
	return ""
}
