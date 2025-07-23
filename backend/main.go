package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		fmt.Fprintln(w, `{"message": "Hello from Go backend!"}`)
	})
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
