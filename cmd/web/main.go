package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/home", Home)
	mux.HandleFunc("/about", About)
	mux.HandleFunc("/blog", Blog)
	mux.HandleFunc("/blog/", BlogPost)
	mux.HandleFunc("/contact", Contact)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}