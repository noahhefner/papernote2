package handlers

import (
	"net/http"
	"fmt"
)

func HandleLogin (w http.ResponseWriter, r *http.Request) {
	fmt.Println("You loggin in bro")
}