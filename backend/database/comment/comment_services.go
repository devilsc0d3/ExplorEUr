package comment

import (
	"errors"
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string
	UserID  int
	PostID  int
}

var comment = &Comment{}

func AddComment(postID int, userID int, message string) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&Comment{PostID: postID, UserID: userID, Message: message})
}

func DeleteComment(commentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Comment{}, commentID)
}

func UpdateCommentMessage(comment string, commentID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Comment{}).Where("id = ?", commentID).Update("comment", comment)
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

func ResetCommentTable() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Migrator().DropTable(&Comment{})
	if err != nil {
		panic("problem to delete comment table")
	}
	err = db.AutoMigrate(&Comment{})
	if err != nil {
		panic("failed to auto migrate: ")
	}
}
