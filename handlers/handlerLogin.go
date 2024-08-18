package handlers

import (
	"html/template"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/pages/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
