package comment

import (
	"errors"
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message    string
	CategoryID int
	UserID     int
	PostID     int
}

func AddComment(message string, postID int, categoryID int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&Comment{Message: message, CategoryID: categoryID, PostID: postID})
}

func DeleteComment(id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Comment{}, id)
}

func Clear() {
	register.Db.Exec("DROP TABLE comments")
}

func UpdateCommentMessage(comment string, id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Comment{}).Where("id = ?", id).Update("comment", comment)
}

func GetComment(comment string) (int, error) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singleComment Comment
	result := db.Select("id").Where("comment = ?", comment).First(&singleComment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic("user not found")
		}
		return 0, result.Error
	}
	return int(singleComment.ID), nil
}
