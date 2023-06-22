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

func DeletePostByUserController(postID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		post.DeletePost(postID)
		return ""
	}
	return "guest mod"
}

func UpdatePostByUserController(content string, postID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		post.UpdatePost(content, postID)
		return ""
	}
	return "guest mod"
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

func DeleteCommentByUserController(commentID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		comment.DeleteComment(commentID)
		return ""
	}
	return "guest mod"
}

func UpdateCommentByUserController(newComment string, commentID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		comment.UpdateCommentMessage(newComment, commentID)
		return ""
	}
	return "guest mod"
}

func AddLikePostByUserController(postID int) string {
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
		return ""
		//idk if it's bullshit or not, need to test
		//like_post.CancelDislikePost(false, postID)
	}
	return "guest mod"
}

func CancelLikePostByUserController(postID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_post.DeleteLikePost(postID)
		return ""
	}
	return "guest mod"

}

func AddDislikePostByUserController(postID int) string {
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
		return ""
		//like_post.CancelLikePost(false, postID)
	}
	return "guest mod"
}

func CancelDislikePostByUserController(postID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_post.DeleteLikePost(postID)
		return ""
	}
	return "guest mod"
}

func AddLikeCommentByUserController(commentID int) string {
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
		return ""
		//idk if it's bullshit or not, need to test
		//like_comment.CancelDislikeComment(false, commentID)
	}
	return "guest mod"
}

func CancelLikeCommentByUserController(commentID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_comment.DeleteLikeComment(commentID)
		return ""
	}
	return "guest mod"
}

func AddDislikeCommentByUserController(commentID int) string {
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
		return ""
		//like_comment.CancelLikeComment(false, commentID)
	}
	return "guest mod"
}

func CancelDislikeCommentByUserController(commentID int) string {
	_, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "" {
		like_comment.DeleteLikeComment(commentID)
		return ""
	}
	return "guest mod"
}

func AskToBeAModerator() {

}
