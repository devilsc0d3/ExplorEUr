package report_comment

import (
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ReportComment struct {
	gorm.Model
	CommentID         int
	NicknameModerator string
	NicknameUser      string
	CommentContent    string
	CategoryID        int
	PostID            int
}

var reportComment = &ReportComment{}

func ReportAComment(commentID int, nicknameModerator string, nicknameUser string, commentContent string, categoryID int, postID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&ReportComment{CommentID: commentID, NicknameModerator: nicknameModerator, NicknameUser: nicknameUser, CommentContent: commentContent, CategoryID: categoryID, PostID: postID})
}
