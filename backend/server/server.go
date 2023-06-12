package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var data = []string{"Place", "Tools", "Information", "+"}

func router() {
	fs := http.FileServer(http.Dir("./front/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/category", category)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/registration", Registration)

	for i := 0; i < len(data); i++ {
		http.HandleFunc("/"+data[i], Chat)
	}
	Reset()
}

const port = "8080"

func Server() {
	router()
	fmt.Println("Listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return
	}

	//log.Fatal(http.ListenAndServeTLS(":"+port, "./localhost.crt", "localhost.key", nil))
}

func Reset() {

	client := &http.Client{}

	// create a new DELETE request
	req, err := http.NewRequest("DELETE", "http://localhost:8080/registration", nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
