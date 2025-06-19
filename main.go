package main

import (
	"fmt"
	asciiart "gomod/asciiart"
	"html/template"
	"net/http"
)

// var temp *template.Template
type Data struct {
	Word  string
	Ascii string
	Style string
}

type Error struct {
	Err  string
	Desc string
}

var error404 = Error{
	Err:  "Error: 404",
	Desc: "Not Found",
}

var error400 = Error{
	Err:  "Error: 400",
	Desc: "Bad Request",
}
var error405 = Error{
	Err:  "Error: 405",
	Desc: "Method Not Allowed",
}
var error500 = Error{
	Err:  "Error: 500",
	Desc: "Internal Server Error",
}

func main() {
	//appending css styles
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css")))) //handling css
	fmt.Println("running server")

	//handler functions
	http.HandleFunc("/", indexHandleFunc)
	http.HandleFunc("/ascii-art", submitHandler)
	// http.HandleFunc("/error", handleError)
	//listener
	http.ListenAndServe(":8000", nil)
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound) // set status code to 404
		executeTemp(w, "error.html", error404)
		//http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	executeTemp(w, "index.html", nil) // execute index template
}

func submitHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		executeTemp(w, "error.html", error404)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		executeTemp(w, "error.html", error405) // if error occurs, execute error template
		return
	}

	word1 := r.FormValue("asciiArt")
	style := r.FormValue("style")
	ascii1, status := asciiart.ProcessASCII(word1, style) //passes word and style

	switch status {
	case 200:
	case 400:
		w.WriteHeader(http.StatusBadRequest)   // set status code to 400
		executeTemp(w, "error.html", error400) // if error occurs, execute error template
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	case 500:
		w.WriteHeader(http.StatusInternalServerError) // set status code to 500
		executeTemp(w, "error.html", error500)        // if error occurs, execute error template
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
	mydata := Data{
		Word:  word1,
		Ascii: ascii1,
		Style: style,
	}
	executeTemp(w, "index.html", mydata) // execute index template with data
}

func executeTemp(w http.ResponseWriter, filename string, data interface{}) {
	tempp, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		if filename != "error.html" {
			w.WriteHeader(http.StatusInternalServerError) // set status code to 500
			executeTemp(w, "error.html", error500)        // if error occurs, execute error template
			return
		}
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	tempp.Execute(w, data)
}
