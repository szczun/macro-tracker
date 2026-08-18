package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("port", ":8000", "HTTP Network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(
		os.Stderr,
		"ERROR\t",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	templateCache, err := newCacheTemplate()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
