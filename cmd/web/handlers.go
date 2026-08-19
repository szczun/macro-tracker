package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/szczun/macro-tracker/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "home.html", nil)
}

func (app *application) userCreate(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "createUser.html", nil)
}

func (app *application) userCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	ageStr := r.PostForm.Get("age")
	weightStr := r.PostForm.Get("weight")
	heightStr := r.PostForm.Get("height")
	activityLevelStr := r.PostForm.Get("activity")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	activityLevel, err := strconv.ParseUint(activityLevelStr, 10, 8)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	user := &models.User{
		Name:          name,
		Age:           age,
		Weight:        weight,
		Height:        height,
		ActivityLevel: uint8(activityLevel),
		TDEE:          2000,
	}

	id, err := app.users.Insert(r.Context(), user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "user id: %v\n%v", id, *user)
}
