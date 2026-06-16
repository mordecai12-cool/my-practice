package main

import (
	"net/http"
)

func okHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ok" {
		w.WriteHeader(http.StatusOK)
		return
	}
}
func badHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/badrequest" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
func notHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/notfound" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
func errorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/error" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/ok", okHandler)
	http.HandleFunc("/badrequest", badHandler)
	http.HandleFunc("/notfound", notHandler)
	http.HandleFunc("/error", errorHandler)
	http.ListenAndServe(":8080", nil)
}
// Test cases:
// curl -i http://localhost:8080/ok
// curl -i http://localhost:8080/badrequest
// curl -i http://localhost:8080/notfound
// curl -i http://localhost:8080/error