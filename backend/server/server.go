package server

import (
	"fmt"
	"log"
	"net/http"
)

func router() {
	fs := http.FileServer(http.Dir("./front/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/category", category)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/registration", Registration)
}

const port = "8080"

func Server() {
	router()
	fmt.Println("Listening on https://localhost:" + port)
	//err := http.ListenAndServe(":"+port, nil)
	//if err != nil {
	//	return
	//}

	log.Fatal(http.ListenAndServeTLS(":"+port, "./localhost.crt", "localhost.key", nil))
}
