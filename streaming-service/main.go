package main

import (
	"fmt"
	"net/http"
)

func streamContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Streaming-Service: Content is streaming!")
}

func main() {
	http.HandleFunc("/stream", streamContentHandler)

	fmt.Println("Streaming-Service is running on port 8082...")
	http.ListenAndServe(":8082", nil)
}

