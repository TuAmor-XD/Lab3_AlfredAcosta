package main

import (
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alfred Acosta"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format(time.RFC1123)))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the Shapes API"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)
	http.HandleFunc("/about", about)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/greeting", greeting)

	http.ListenAndServe(":4000", nil)
}

