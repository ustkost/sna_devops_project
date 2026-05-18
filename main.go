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

func load(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100000000; i++ {
		_ = i * i
	}

	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "CPU loaded on pod %s\n", hostname)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)
	http.HandleFunc("/load", load)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on :8080")
}
