package report_post

import "gorm.io/gorm"

type ReportPost struct {
	gorm.Model
	PostID            string
	NicknameModerator string
	NicknameUser      string
	PostContent       string
	CategoryID        int
}

var reportPost = &ReportPost{}

func ReportAPost() {

}
