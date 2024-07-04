package main

import (
	"fmt"
	"net/http"
	"nhefner/papernote2/handlers"
	"nhefner/papernote2/atlasconfig"
)

func main () {

	err := atlasconfig.ReadATLASConfig()
	if err != nil {
		panic("Failed to read config file")
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /login", handlers.HandleLogin)
	router.HandleFunc("GET /home", handlers.HandleHome)

	fmt.Println("Running server at localhost:8080")

	err = http.ListenAndServe(":8080", router)

	if err != nil {
		panic("Failed to start server")
	}

}