package main

import (
	"fmt"
	"net/http"
)

func processPaymentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Payment-Service: Payment processed!")
}

func main() {
	http.HandleFunc("/process", processPaymentHandler)

	fmt.Println("Payment-Service is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}

