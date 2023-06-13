package server

import (
	"fmt"
	"net/http"
)

const port = "8080"

func Router() {
	fs := http.FileServer(http.Dir("./front/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/category", CategoryHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/registration", RegistrationHandler)
	http.HandleFunc("/recovering_password", RecoveringPassword)
}

func Server() {
	Router()
	fmt.Println("Listening on https://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}
}
