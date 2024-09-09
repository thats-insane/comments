package main

import (
	"fmt"
	"net/http"
)

func (a *appDependencies) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: available")
	fmt.Fprintf(w, "Env: %s\n", a.config.env)
	fmt.Fprintf(w, "Version: %s\n", appVersion)
}
