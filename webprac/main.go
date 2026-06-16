package main

import (
	"html/template"
	"net/http"
	"fmt"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		templ, err := template.ParseFiles("templates/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "fail to load page")
			return
		}
		templ.Execute(w, nil)
	}
}
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	text := r.FormValue("input")
	fmt.Fprintln(w, text)
}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/echo", echoHandler)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}