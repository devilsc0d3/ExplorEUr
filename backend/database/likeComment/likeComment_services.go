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

func AddLikeComment(isLike bool, isDislike bool) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&LikeComment{isLike: isLike, isDislike: isDislike})
}

func DeleteLikeComment(id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&LikeComment{}, id)
}

func UpdateLikeCommentIsLike(isLike bool, id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikeComment{}).Where("id = ?", id).Update("isLike", isLike)
}

func UpdateLikeCommentIsDislike(isDislike bool, id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&LikeComment{}).Where("id = ?", id).Update("isDislike", isDislike)
}
