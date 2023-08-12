package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/detail", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.Handle("/static/", http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./ui/static/"))))

	log.Print("Starting server on 4000 port")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
