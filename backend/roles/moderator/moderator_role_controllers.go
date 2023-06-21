package moderator

import (
	"exploreur/backend/register"
	"exploreur/backend/report_comment"
	"exploreur/backend/report_post"
)

func ReportPostByModeratorController(postID int, nicknameUser string, postContent string, categoryID int) string {
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "moderator" {
		return "not moderator mod"
	}
	report_post.ReportAPost(postID, nickname, nicknameUser, postContent, categoryID)
	return ""
}

func ReportCommentByModeratorController(commentID int, nicknameUser string, commentContent string, categoryID int, postID int) string {
	nickname, role, err := register.DecodeJWTToken(register.Token)
	if err != nil {
		panic("token error")
	}
	if role != "admin" {
		return "not admin mod"
	}
	report_comment.ReportAComment(commentID, nickname, nicknameUser, commentContent, categoryID, postID)
	return ""
}
