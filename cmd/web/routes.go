package main

import "net/http"

func (app *application) routes(staticDir string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/detail", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
