package main

import (
	"net/http"
	asciiart "gomod/asciiart"
)

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
