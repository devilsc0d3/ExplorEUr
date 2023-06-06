package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string
	Email    string
	Password string
	PostID   int
	Role     string
}

var user = &User{}

func AddUser() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&User{Nickname: user.Nickname, Email: user.Email, Password: user.Password})
}

func DeleteUser(id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Delete(&User{}, id)
}

func UpdateUserNickname(nickname string) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Save(&User{Nickname: nickname})
}

func UpdateUserPassword(password string) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Save(&User{Password: password})
}

func UpdateUserRole(role string) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Save(&User{Role: role})
}
