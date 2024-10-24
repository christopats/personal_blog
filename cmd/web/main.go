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
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))


	log.Println("Starting server on :42069")
	err := http.ListenAndServe(":42069", mux)
	log.Fatal(err)

}