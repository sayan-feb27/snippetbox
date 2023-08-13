package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

const LogFormat = log.LUTC | log.Ldate | log.Ltime | log.Lshortfile

func main() {
	addr := flag.String("addr", ":4000", "HTTP network addresses")
	staticDir := flag.String("static_dir", "./ui/static/", "HTTP network addresses")

	flag.Parse()

	app := &application{
		errorLog: log.New(os.Stdout, "[ERROR]\t", LogFormat),
		infoLog:  log.New(os.Stdout, "[INFO]\t", LogFormat),
	}
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(*staticDir),
	}

	app.infoLog.Printf("Starting server on %s port", *addr)

	err := server.ListenAndServe()
	app.errorLog.Fatal(err)
}
