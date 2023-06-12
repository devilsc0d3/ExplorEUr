package userDB

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetEnv(key string) string {
	value, isValid := os.LookupEnv(key)
	if !isValid {
		panic("Not found")
	}
	return value
}

func Init() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
