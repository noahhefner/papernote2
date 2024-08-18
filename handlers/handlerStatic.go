package handlers

import (
	"net/http"
)

func HandleLogo(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "static/png/logo.png")

}
