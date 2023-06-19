package comment

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

var Db *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Comment{})
	if err != nil {
		return
	}
	Db = db
}
