package main

import "net/http"

const port = "8765"

func main() {
	router()
	http.ListenAndServe(port, nil)
	//TODO BDD
}
