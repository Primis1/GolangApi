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

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO: \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: \t", log.Ltime|log.Llongfile)


	app := &application{
		errorLog: errorLog, 
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// errorLog := log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Printf("Starting server on %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
