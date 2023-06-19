package post

import (
	"errors"
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content    string
	UserID     int
	CategoryID int
}

func AddPost(content string, userID int, categoryID int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&Post{Content: content, UserID: userID, CategoryID: categoryID})
}

func DeletePost(id int) {
	db, err := gorm.Open(postgres.Open(register.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Post{}, id)
}

func Clear() {
	register.Db.Exec("DROP TABLE posts")
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
