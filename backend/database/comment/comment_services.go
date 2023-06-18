package comment

import (
	"errors"
	"exploreur/backend/register"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string
}

func AddComment(message string) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("tst1")

	db.Create(&Comment{Message: message})
}

func DeleteComment(id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Comment{}, id)
}

func UpdateCommentMessage(comment string, id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Comment{}).Where("id = ?", id).Update("comment", comment)
}

func GetComment(comment string) (int, error) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
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
