package main

import (
	"encoding/json"
	"net/http"
)

func (a *appDependencies) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": a.config.env,
		"version":     appVersion,
	}

	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, "The server encountered an issue and was not able to process your request", http.StatusInternalServerError)
		return
	}
}

func (a *appDependencies) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	jsResponse, err := json.Marshal(data)
	if err != nil {
		return err
	}

	jsResponse = append(jsResponse, '\n')
	for key, value := range headers {
		w.Header().Set(key, value)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsResponse)

	return nil
}
