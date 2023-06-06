package main

import "net/http"

const port = "8765"

func main() {
	server()
	//TODO BDD
}

func server() {
	router()
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}
}
