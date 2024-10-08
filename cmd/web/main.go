package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/blog", Home)
	mux.HandleFunc("/about", Home)
	mux.HandleFunc("/contact", Home)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}