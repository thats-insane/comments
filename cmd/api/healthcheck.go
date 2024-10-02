package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

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
