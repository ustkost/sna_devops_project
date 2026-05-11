package main

import (
	"fmt"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()

	fmt.Fprintf(w, "Hello, Kubernetes!\n")
	fmt.Fprintf(w, "Pod: %s\n", hostname)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
