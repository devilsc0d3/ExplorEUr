package server

import (
	"exploreur/backend/register"
	"fmt"
	"net/http"
	"strconv"
)

var CategoriesId []int
var RegisteredPaths = make(map[int]bool) // Map to track registered paths

func Router() {
	fs := http.FileServer(http.Dir("./front/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/category", CategoryHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/registration", RegistrationHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/easter_egg", EasterEgg)
	http.HandleFunc("/info", Info)
	http.HandleFunc("/activity", ActivityHandler)
	http.HandleFunc("/recover_password", RecoverHandler)

	register.Db.Table("categories").Pluck("id", &CategoriesId)

	for i := 0; i < len(CategoriesId); i++ {
		http.HandleFunc("/"+strconv.Itoa(CategoriesId[i]), Chat)
		RegisteredPaths[CategoriesId[i]] = true
	}

}

const port = "8080"

func Server() {

	Router()
	fmt.Println("Listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}

	//log.Fatal(http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", nil))
}

//func AddRouteCategory() {
//	for i := 0; i < len(CategoriesId); i++ {
//		if !RegisteredPaths[CategoriesId[i]] { // Check if path is already registered
//			http.HandleFunc("/"+strconv.Itoa(CategoriesId[i]), Chat)
//			RegisteredPaths[CategoriesId[i]] = true // Mark path as registered
//		}
//	}
//}
