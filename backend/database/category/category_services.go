package category

import (
	"errors"
	"exploreur/backend/register"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
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
	var id []int
	Db.Table("categories").Order("created_at DESC").Pluck("id", &id)
	//AddRouteCategory(id[len(id)-1])
	//// Todo divide page
	//// Todo divide bdd
}

func HomeHandler2(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")

	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func AddRouteCategory(CategoriesId int) {
	http.HandleFunc("/"+strconv.Itoa(CategoriesId+1), HomeHandler2)
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
