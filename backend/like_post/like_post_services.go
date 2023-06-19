package like_post

import (
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LikePost struct {
	gorm.Model
	IsLike    bool
	IsDislike bool
	UserID    int
	PostID    int
}

var likePost = &LikePost{}

func AddLikePost(isLike bool, isDislike bool, userID int, postID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&LikePost{IsLike: isLike, IsDislike: isDislike, UserID: userID, PostID: postID})
}

func DeleteLikePost(likePostID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&LikePost{}, likePost)
}

func CancelLikePost(isLike bool, likePostID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikePost{}).Where("id = ?", likePostID).Update("isLike", isLike)
}

func CancelDislikePost(isDislike bool, likePostID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikePost{}).Where("id = ?", likePostID).Update("isDislike", isDislike)
}

func ResetLikePostTable() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Migrator().DropTable(&LikePost{})
	if err != nil {
		panic("problem to delete likePost table")
	}
	err = db.AutoMigrate(&LikePost{})
	if err != nil {
		panic("failed to auto migrate: ")
	}
}
