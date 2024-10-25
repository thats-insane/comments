package main

import (
	"fmt"
	"net/http"

	"github.com/thats-insane/comments/internal/data"
	"github.com/thats-insane/comments/internal/validator"
)

func (a *appDependencies) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	comment := &data.Comment{
		Content: incomingData.Content,
		Author:  incomingData.Author,
	}

	v := validator.New()

	data.ValidateComment(v, comment)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", incomingData)

}
