package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "home.html")
}

func (app *application) userCreate(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "createUser.html")
}
