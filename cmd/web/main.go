package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/szczun/macro-tracker/internal/models"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	users         models.UserRepository
}

func main() {
	addr := flag.String("port", ":8000", "HTTP Network address")
	dsn := flag.String(
		"dsn",
		"macro_user:fit4tu/macro_tracker?parseTime=true",
		"MySQL dns",
	)

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(
		os.Stderr,
		"ERROR\t",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	templateCache, err := newCacheTemplate()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
		users:         models.NewUserModel(db),
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
