package post

import (
	"errors"
	"exploreur/backend/register"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string
	UserID  int
}

var post = &Post{}

func AddPost(content string) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("test1")
	db.Create(&Post{Content: content})
	fmt.Println("test2")
}

func DeletePost(id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Post{}, id)
}

func UpdatePost(content string, id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Post{}).Where("id = ?", id).Update("content", content)
}

func GetPost(content string) (int, error) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singlePost Post
	result := db.Select("id").Where("content = ?", content).First(&singlePost)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic("post not found")
		}
		return 0, result.Error
	}
	return int(singlePost.ID), nil
}
