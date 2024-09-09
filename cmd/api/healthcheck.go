package main

import (
	"fmt"
	"net/http"
)

func (a *appDependencies) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	jsResponse := `{"status": "available", "environment": %q, "version": %q}`

	jsResponse = fmt.Sprintf(jsResponse, a.config.env, appVersion)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsResponse))
	// fmt.Fprintln(w, "Status: available")
	// fmt.Fprintf(w, "Env: %s\n", a.config.env)
	// fmt.Fprintf(w, "Version: %s\n", appVersion)
}
