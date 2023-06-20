package report_comment

import (
	"gorm.io/gorm"
)

type ReportComment struct {
	gorm.Model
	CommentID         string
	NicknameModerator string
	NicknameUser      string
	CommentContent    string
	CategoryID        int
	PostID            int
}

var reportComment = &ReportComment{}

func ReportAComment() {

}
