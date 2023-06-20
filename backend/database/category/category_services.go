package category

import (
	"errors"
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string
	PostID int
}

func AddCategory(name string) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&Category{Name: name})
	//server.AddRouteCategory()

}

func DeleteCategory(id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&Category{}, id)
}

func Clear() {
	register.Db.Exec("DROP TABLE categories")
}

func UpdateCategoryName(name string, id int) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Category{}).Where("id = ?", id).Update("name", name)
}

func GetCategory(name string) (int, error) {
	db, err := gorm.Open(postgres.Open(GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var singleCategory Category
	result := db.Select("id").Where("name = ?", name).First(&singleCategory)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic("user not found")
		}
		return 0, result.Error
	}
	return int(singleCategory.ID), nil
}
