package server

import (
	"fmt"
	"log"
	"net/http"
)

func router() {
	fs := http.FileServer(http.Dir("../front/static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/category", category)
}

const port = "8080"

func Server() {
	router()
	fmt.Println("Listening on http://localhost:" + port)
	fmt.Println("Listening on https://localhost:" + port)
	//err := http.ListenAndServe(":"+port, nil)
	//if err != nil {
	//	return
	//}

	//log.Fatal(http.ListenAndServeTLS(":"+port, "./localhost.crt", "localhost.key", nil))

	err := http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil)
	log.Fatal(err)
}
