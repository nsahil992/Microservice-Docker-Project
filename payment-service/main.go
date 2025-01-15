package main

import (
	"encoding/json"
	"net/http"
)

type PaymentRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	json.NewDecoder(r.Body).Decode(&req)

	// Simulate payment processing
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment successful"))
}

func main() {
	http.HandleFunc("/pay", paymentHandler)
	http.ListenAndServe(":8082", nil)
}

