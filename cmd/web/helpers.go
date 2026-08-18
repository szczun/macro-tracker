package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (app *application) render(w http.ResponseWriter, status int, page string) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template does not exist")
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", "")
	if err != nil {
		app.serverError(w, err)
	}
}
