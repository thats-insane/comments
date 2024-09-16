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

	jsResponse, err := json.Marshal(data)
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, "The server encountered an issue and was not able to process your request", http.StatusInternalServerError)
		return
	}

	jsResponse = append(jsResponse, '\n')
	// jsResponse := `{"status": "available", "environment": %q, "version": %q}`

	// jsResponse = fmt.Sprintf(jsResponse, a.config.env, appVersion)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsResponse)
	// fmt.Fprintln(w, "Status: available")
	// fmt.Fprintf(w, "Env: %s\n", a.config.env)
	// fmt.Fprintf(w, "Version: %s\n", appVersion)
}
