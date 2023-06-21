package user

import (
	"exploreur/backend/database/comment"
	"exploreur/backend/database/like_comment"
	"exploreur/backend/like_post"
	"exploreur/backend/post"
	"exploreur/backend/register"
)

func AddPostByUserController(input string, categoryID int) string {

	input = FilterInsults2(input)

	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		id, err := register.GetIDByNickname(nickname)
		if err != nil {
			panic("GetIDByNickname error")
		}
		post.AddPost(input, id, categoryID)
		return ""
	}
	return "guest mod"
}

func DeletePostByUserController(postID int) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "user" || role == "moderator" {
		post.DeletePost(postID)
	}
}

func UpdatePostByUserController() {
}

func AddCommentByUserController(postID int, input string, categoryID int) string {
	if !CheckInsults(input) {
		return "there is at least one insult in the text"
	}
	if !CheckLength(input) {
		return "the text are too long"
	}
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		id, err := register.GetIDByNickname(nickname)
		if err != nil {
			panic("GetIDByNickname error")
		}
		comment.AddComment(input, categoryID, id, postID)
		return ""
	}
	return "guest mod"
}

func DeleteCommentByUserController(commentID int) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role == "user" || role == "moderator" {
		comment.DeleteComment(commentID)
	}
}

func UpdateCommentByUserController() {
}

func AddLikePostByUserController(postID int) {
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		id, err := register.GetIDByNickname(nickname)
		if err != nil {
			panic("GetIDByNickname error")
		}
		like_post.AddLikePost(true, false, id, postID)
		//idk if it's bullshit or not, need to test
		//like_post.CancelDislikePost(false, postID)
	}
}

func CancelLikePostByUserController(postID int) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_post.DeleteLikePost(postID)
	}
}

func AddDislikePostByUserController(postID int) {
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		id, err := register.GetIDByNickname(nickname)
		if err != nil {
			panic("GetIDByNickname error")
		}
		like_post.AddLikePost(false, true, id, postID)
		//like_post.CancelLikePost(false, postID)
	}
}

func CancelDislikePostByUserController(postID int) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_post.DeleteLikePost(postID)
	}
}

func AddLikeCommentByUserController(commentID int) {
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		id, err := register.GetIDByNickname(nickname)
		if err != nil {
			panic("GetIDByNickname error")
		}
		like_comment.AddLikeComment(true, false, id, commentID)
		//idk if it's bullshit or not, need to test
		//like_comment.CancelDislikeComment(false, commentID)
	}
}

func CancelLikeCommentByUserController(commentID int) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_comment.DeleteLikeComment(commentID)
	}
}

func AddDislikeCommentByUserController(commentID int) {
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		id, err := register.GetIDByNickname(nickname)
		if err != nil {
			panic("GetIDByNickname error")
		}
		like_comment.AddLikeComment(false, true, id, commentID)
		//like_comment.CancelLikeComment(false, commentID)
	}
}

func CancelDislikeCommentByUserController(commentID int) {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_comment.DeleteLikeComment(commentID)
	}
}

func AskToBeAModerator() {

}
