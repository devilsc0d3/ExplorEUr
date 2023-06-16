package register

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string
	Email    string
	Password []byte
	PostID   int
	Role     string
}

var user = &User{}

func AddUser(nickname string, email string, password string, role string) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
	db.Create(&User{Nickname: nickname, Email: email, Password: passwordHashed, Role: role})
}

func DeleteUser(id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&User{}, id)
}

func UpdateUserNickname(nickname string, id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&User{}).Where("id = ?", id).Update("nickname", nickname)
}

func UpdateUserPassword(password string, id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&User{}).Where("id = ?", id).Update("password", password)
}
func UpdateUserRole(role string, id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&User{}).Where("id = ?", id).Update("role", role)
}

func GetIDByNickname(nickname string) (int, error) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singleUser User
	result := db.Select("id").Where("nickname = ?", nickname).First(&singleUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return -1, nil
		}
		return -1, result.Error
	}
	return int(singleUser.ID), nil
}

func GetIDByEmail(email string) (int, error) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singleUser User
	result := db.Select("id").Where("email = ?", email).First(&singleUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return -1, nil
		}
		return -1, result.Error
	}
	return int(singleUser.ID), nil
}

func GetUserByID(id int) *User {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return nil
	}

	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}

	return &user
}

func GetIDByUser(user *User) int {
	return int(user.ID)
}

func ResetDatabase() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Migrator().DropTable(&User{})
	if err != nil {
		panic("problem to delete user table")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to auto migrate: ")
	}
}
