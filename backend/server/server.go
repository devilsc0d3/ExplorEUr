package server

import (
	"fmt"
	"net/http"
)

const port = "8080"

func Server() {
	Router()
	fmt.Println("Listening on https://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}
}
