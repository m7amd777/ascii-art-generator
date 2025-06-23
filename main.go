package main

import (
	"fmt"
	"net/http"
)


func main() {
	//appending css styles
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css")))) //handling css
	fmt.Println("running server at http://localhost:8000")

	//handler functions
	http.HandleFunc("/", indexHandleFunc)
	http.HandleFunc("/ascii-art", submitHandler)
	http.ListenAndServe(":8000", nil)
}
