package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {

	var fileNames []string

	files, err := ioutil.ReadDir("./notes")
	if err != nil {
		http.Error(w, "Failed to read notes.", http.StatusInternalServerError)
		return
	}

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	t, err := template.ParseFiles("templates/pages/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, fileNames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
