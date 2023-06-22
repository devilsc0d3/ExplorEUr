package like_comment

import (
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LikeComment struct {
	gorm.Model
	IsLike    bool
	IsDislike bool
	UserID    int
	CommentID int
}

func AddLikeComment(isLike bool, isDislike bool, userID int, commentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&LikeComment{IsLike: isLike, IsDislike: isDislike, UserID: userID, CommentID: commentID})
}

func DeleteLikeComment(likeCommentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&LikeComment{}, likeCommentID)
}

func CancelLikeComment(isLike bool, likeCommentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikeComment{}).Where("id = ?", likeCommentID).Update("isLike", isLike)
}

func CancelDislikeComment(isDislike bool, likeCommentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikeComment{}).Where("id = ?", likeCommentID).Update("isDislike", isDislike)
}

func ResetLikeCommentTable() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Migrator().DropTable(&LikeComment{})
	if err != nil {
		panic("problem to delete like_comment table")
	}
	err = db.AutoMigrate(&LikeComment{})
	if err != nil {
		panic("failed to auto migrate: ")
	}
}
