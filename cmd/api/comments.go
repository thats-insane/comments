package main

import (
	"fmt"
	"net/http"

	_ "github.com/thats-insane/comments/internal/data"
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

	fmt.Fprintf(w, "%+v\n", incomingData)
}
