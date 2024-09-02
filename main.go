package main

import (
	"fmt"
	"net/http"
	"nhefner/papernote2/handlers"
)

func main() {

	router := http.NewServeMux()

	// pages
	router.HandleFunc("GET /login", handlers.HandleLogin)
	router.HandleFunc("GET /home", handlers.HandleHome)
	router.HandleFunc("GET /note/editor", handlers.HandleGetEditor)
	router.HandleFunc("GET /note/view", handlers.HandleGetRendered)

	// static files
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Running server at localhost:8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic("Failed to start server")
	}

}
