package server

import (
	"fmt"
	"net/http"
	"strconv"
)

var data = map[int]string{0: "Place", 1: "Tools", 2: "Information", 3: "+"}
var registeredPaths = make(map[int]bool) // Map to track registered paths

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

	for id := range data {
		http.HandleFunc("/"+strconv.Itoa(id), Chat)
		registeredPaths[id] = true // Mark path as registered
	}
	//Reset()

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

func Reset() {

	ID := 5
	categoryTitle := ""
	data[ID] = categoryTitle
	for id := range data {
		if !registeredPaths[id] { // Check if path is already registered
			http.HandleFunc("/"+strconv.Itoa(id), Chat)
			registeredPaths[id] = true // Mark path as registered
		}
	}
}
