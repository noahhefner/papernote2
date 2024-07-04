package handlers

import (
	"net/http"
	"html/template"
	"nhefner/papernote2/atlasconfig"
)

func HandleHome (w http.ResponseWriter, r *http.Request) {
		
	t, err := template.ParseFiles("templates/pages/home.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, atlasconfig.Config)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}