package register

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetEnv(key string) string {
	env := os.Getenv(key)
	return env
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
