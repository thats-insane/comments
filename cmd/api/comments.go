package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/thats-insane/comments/internal/data"
)

func (a *appDependencies) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	err := json.NewDecoder(r.Body).Decode(&incomingData)
	if err != nil {
		a.errResponseJSON(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", incomingData)
}
