package main

import (
	"errors"
	"net/http"

	"github.com/thats-insane/comments/internal/data"
	"github.com/thats-insane/comments/internal/validator"
)

func (a *appDependencies) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := a.readJSON(w, r, incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Username: incomingData.Username,
		Email: incomingData.Email,
		Activated: false,
	}

	err = user.Password.Set(incomingData.Password)

	if err != nil package main
	
	func main() {
		a.serverErrResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateUser(v, user)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.UserModel.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email already exists")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrResponse(w, r, err)
		}
		return
	}

	data := envelope{
		"user": user,
	}

	err = a.writeJSON(w, http.StatusCreated, data, nil)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}	
}