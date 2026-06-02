package main

import (
	"fmt"
	"net/http"
)

func job(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
func main() {
	http.HandleFunc("/marvy", job)
	fmt.Println("loading server at: http://localhost:8080/marvy")
	http.ListenAndServe(":8080", nil)
}
