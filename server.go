package main

import "net/http"

func router() {
	http.HandleFunc("/", test)
}

func test(w http.ResponseWriter, _ *http.Request) {

}
