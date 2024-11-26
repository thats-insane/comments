package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *appDependencies) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(a.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(a.notAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/comments", a.listCommentsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/comments/:id", a.displayCommentHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/comments/:id", a.updateCommentHandler)

	router.HandlerFunc(http.MethodPost, "/v1/comments", a.createCommentHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", a.registerUserHandler)

	router.HandlerFunc(http.MethodDelete, "/v1/comments/:id", a.deleteCommentHandler)

	return a.recoverPanic(router)
}
