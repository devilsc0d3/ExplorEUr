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
	fmt.Println("Listening on https://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}

	//log.Fatal(http.ListenAndServeTLS(":"+port, "localhost.crt", "localhost.key", nil))
}
