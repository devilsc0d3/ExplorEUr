package server

import (
	"fmt"
	"net/http"
	"strings"
)

var data = []string{"Place", "Tools", "Information", "+"}
var registeredPaths = make(map[string]bool) // Map to track registered paths

func router() {
	fs := http.FileServer(http.Dir("./front/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/category", category)

	for i := 0; i < len(data); i++ {
		http.HandleFunc("/"+strings.ToLower(data[i]), Chat)
		registeredPaths[data[i]] = true // Mark path as registered
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

	data = append(data, "test")
	for i := 0; i < len(data); i++ {
		if !registeredPaths[data[i]] { // Check if path is already registered
			http.HandleFunc("/"+strings.ToLower(data[i]), Chat)
			registeredPaths[data[i]] = true // Mark path as registered
		}
	}
}
