package main

import (
	"fmt"
	"net/http"
)

const port = "8080"

func main() {
	server()
	//TODO BDD
}

func server() {
	router()
	fmt.Println("Listening on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return
	}
}
