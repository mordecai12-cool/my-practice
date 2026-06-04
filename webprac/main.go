package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.Error(w, "Page Not Fond", http.StatusNotFound)
	}

	
	if r.Method != http.MethodGet{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
	tem, err := template.ParseFiles("template/index.html")
	if err != nil{
		http.Error(w, "Sever Error", http.StatusInternalServerError)
	}
	tem.Execute(w, nil)
}
func asciiHandler(w http.ResponseWriter, r *http.Request){
	term, err := template.ParseFiles("template/index.html")
	if err != nil{
		http.Error(w, "internal marvelous error", http.StatusInternalServerError)
	}
	r.FormValue("text")
	term.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
