package server

import (
	"fmt"
	"net/http"
	"strconv"
)

var data = map[int]string{0: "Place", 1: "Tools", 2: "Information", 3: "+"}
var registeredPaths = make(map[int]bool) // Map to track registered paths

func router() {
	fs := http.FileServer(http.Dir("./front/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/category", category)
	http.HandleFunc("/info", Info)

	for id := range data {
		http.HandleFunc("/"+strconv.Itoa(id), Chat)
		registeredPaths[id] = true // Mark path as registered
	}
	//Reset()

}

const port = "8080"

func Server() {
	router()
	fmt.Println("Listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}

	//log.Fatal(http.ListenAndServeTLS(":"+port, "./localhost.crt", "localhost.key", nil))
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
