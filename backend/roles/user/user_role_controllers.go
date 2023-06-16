package user

import (
	"exploreur/backend/register"
	"fmt"
)

func AddPostByUserController(input string) string {
	if !CheckInsults(input) {
		return "there is at least one insult in the text"
	}
	if !CheckLength(input) {
		return "the text are too long"
	}
	fmt.Println(register.DecodeJWTToken(register.Token))
	//fmt.Printf("NICKNAME:", nickname)
	//fmt.Printf("ROLE:", role)
	//id, _ := register.GetIDByNickname(nickname)
	//post.AddPost(id, input)
	return ""
}
