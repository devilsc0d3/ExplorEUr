package server

import (
	"fmt"
	"net/http"
)

func router() {
	http.HandleFunc("/", home)
	http.HandleFunc("/category", category)
}

const port = "8080"

func Server() {
	router()
	fmt.Println("Listening on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return
	}
}
