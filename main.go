package main

import (
	"fmt"
	"net/http"
	"nhefner/papernote2/handlers"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("GET /login", handlers.HandleLogin)
	router.HandleFunc("GET /home", handlers.HandleHome)

	router.HandleFunc("GET /static/png/logo.png", handlers.HandleLogo)

	fmt.Println("Running server at localhost:8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic("Failed to start server")
	}

}
