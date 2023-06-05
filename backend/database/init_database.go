package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type User struct {
	gorm.Model
	Nickname  string
	Email     string
	Password  string
	CountryID int
	PostID    int
}

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

	db.AutoMigrate(&User{})

	db.Create(&User{Nickname: "admin", Email: "admin@forum.com", Password: "exploreur"})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}
