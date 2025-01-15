package main

import (
	"net/http"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Streaming content..."))
}

func main() {
	http.HandleFunc("/stream", streamHandler)
	http.ListenAndServe(":8083", nil)
}

