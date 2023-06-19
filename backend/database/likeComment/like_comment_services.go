package likeComment

import (
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LikeComment struct {
	gorm.Model
	isLike    bool
	isDislike bool
}

var likeComment = &LikeComment{}

func AddLikeComment(isLike bool, isDislike bool) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&LikeComment{isLike: isLike, isDislike: isDislike})
}

func DeleteLikeComment(likeCommentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&LikeComment{}, likeCommentID)
}

func CancelLikeCommentIsLike(isLike bool, likeCommentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikeComment{}).Where("id = ?", likeCommentID).Update("isLike", isLike)
}

func CancelLikeCommentIsDislike(isDislike bool, likeCommentID int) {
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
		panic("problem to delete likeComment table")
	}
	err = db.AutoMigrate(&LikeComment{})
	if err != nil {
		panic("failed to auto migrate: ")
	}
}
