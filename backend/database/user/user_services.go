package user

import (
	"errors"
	"exploreur/backend/database"
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

func AddUser(nickname string, email string, password string, role string) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&User{Nickname: nickname, Email: email, Password: password, Role: role})
}

func DeleteUser(id int) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&User{}, id)
}

func UpdateUserNickname(nickname string, id int) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&User{}).Where("id = ?", id).Update("nickname", nickname)
}

func UpdateUserPassword(password string, id int) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&User{}).Where("id = ?", id).Update("password", password)
}

func UpdateUserRole(role string, id int) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&User{}).Where("id = ?", id).Update("role", role)
}

func GetUser(nickname string) (int, error) {
	db, err := gorm.Open(postgres.Open(database.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singleUser User
	result := db.Select("id").Where("nickname = ?", nickname).First(&singleUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic("user not found")
		}
		return 0, result.Error
	}
	return int(singleUser.ID), nil
}
