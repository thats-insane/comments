package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type envelope map[string]any

func (a *appDependencies) readJSON(w http.ResponseWriter, r *http.Request, destination any) error {
	err := json.NewDecoder(r.Body).Decode(destination)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError
		var invalidUnmarshalErr *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxErr):
			return fmt.Errorf("the body contains badly-formed JSON at character %d", syntaxErr.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("the body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeErr):
			if unmarshalTypeErr.Field != "" {
				return fmt.Errorf("the body contains the incorrect JSON type for field %q", unmarshalTypeErr.Field)
			}
			return fmt.Errorf("the body contains the incorrect JSON type (at character %d)", unmarshalTypeErr.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("the body must not be empty")
		case errors.As(err, &invalidUnmarshalErr):
			panic(err)
		default:
			return err
		}
	}

	return nil
}

func (a *appDependencies) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.config.env,
			"version":     appVersion,
		},
	}

	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrResponse(w, r, err)
	}
}

func (a *appDependencies) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	jsResponse, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	jsResponse = append(jsResponse, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(jsResponse)
	if err != nil {
		return err
	}

	return nil
}
