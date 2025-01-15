package main

import (
	"fmt"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple welcome message
	fmt.Fprintln(w, "Welcome to the login page")
}

func main() {
	http.HandleFunc("/", welcomeHandler) // Root path for the service
	fmt.Println("Auth Service is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}

