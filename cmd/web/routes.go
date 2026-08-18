package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	fs := http.FileServer(http.Dir("./ui/static/"))
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", app.home)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("GET /create/user", app.userCreate)
	//mux.HandleFunc("POST /user/create", app.userCreatePost)

	return app.serverRecover(secureHeaders(mux))
}
