package main

import (
	"fmt"
	"net/http"
	"html/template"
)


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