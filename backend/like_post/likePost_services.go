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

func AddLikePost(isLike bool, isDislike bool) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&LikePost{IsLike: isLike, IsDislike: isDislike})
}

func DeleteLikePost(id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&LikePost{}, id)
}

func UpdateLikePostIsLike(isLike bool, id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikePost{}).Where("id = ?", id).Update("isLike", isLike)
}

func UpdateLikePostIsDislike(isDislike bool, id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikePost{}).Where("id = ?", id).Update("isDislike", isDislike)
}
