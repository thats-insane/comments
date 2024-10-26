package main

import (
	"errors"
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

	err = a.commentModel.Insert(comment)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/comments/%d", comment.ID))

	data := envelope{
		"comment": comment,
	}

	err = a.writeJSON(w, http.StatusCreated, data, headers)

	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", incomingData)
}

func (a *appDependencies) displayCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)

	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	comment, err := a.commentModel.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrResponse(w, r, err)
		}

		return
	}

	data := envelope{
		"comment": comment,
	}

	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}
}

func (a *appDependencies) updateCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)

	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	comment, err := a.commentModel.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrResponse(w, r, err)
		}

		return
	}

	var incomingData struct {
		Content *string `json:"content"`
		Author  *string `json:"author"`
	}

	err = a.readJSON(w, r, &incomingData)

	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	if incomingData.Content != nil {
		comment.Content = *incomingData.Content
	}

	err = a.commentModel.Update(comment)

	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	data := envelope{
		"comment": comment,
	}

	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}
}

func (a *appDependencies) deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)

	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	err = a.commentModel.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrResponse(w, r, err)
		}

		return
	}

	data := envelope{
		"message": "comment successfully deleted",
	}

	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrResponse(w, r, err)
	}
}

func (a *appDependencies) listCommentsHandler(w http.ResponseWriter, r *http.Request) {
	var queryParametersData struct {
		Content string
		Author  string
	}

	queryParameters := r.URL.Query()
	queryParametersData.Content = a.getSingleQueryParameters(queryParameters, "content", "")
	queryParametersData.Author = a.getSingleQueryParameters(queryParameters, "author", "")

	comments, err := a.commentModel.GetAll(queryParametersData.Content, queryParametersData.Author)

	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	data := envelope{
		"comments": comments,
	}

	err = a.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		a.serverErrResponse(w, r, err)
	}
}
