package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Auth-Service: Login successful!")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Auth-Service: Signup successful!")
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)

	fmt.Println("Auth-Service is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

