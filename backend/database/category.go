package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Nickname string
	Email    string
	Password string
	PostID   int
}

var category = &Category{}

func AddCategory() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&Category{Nickname: user.Nickname, Email: user.Email, Password: user.Password})
}

func SuppCategory(id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Delete(&Category{}, id)

}
