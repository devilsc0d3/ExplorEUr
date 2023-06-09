package report_post

import (
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ReportPost struct {
	gorm.Model
	PostID            int
	NicknameModerator string
	NicknameUser      string
	PostContent       string
	CategoryID        int
}

var reportPost = &ReportPost{}

func ReportAPost(postID int, nicknameModerator string, nicknameUser string, postContent string, categoryID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&ReportPost{PostID: postID, NicknameModerator: nicknameModerator, NicknameUser: nicknameUser, PostContent: postContent, CategoryID: categoryID})
}

func Clear() {
	register.Db.Exec("DROP TABLE report_posts")
}
