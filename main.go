package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current Date & Time: %s", currentTime)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
