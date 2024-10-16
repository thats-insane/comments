package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq v1.10.9"
)

const appVersion = "1.0.0"

type serverConfig struct {
	port int
	env  string
}

type appDependencies struct {
	config serverConfig
	logger *slog.Logger
}

func main() {
	var settings serverConfig

	flag.IntVar(&settings.port, "port", 4000, "Server Port")
	flag.StringVar(&settings.env, "env", "development", "Environment(Development|Staging|Production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	appInstance := &appDependencies{
		config: settings,
		logger: logger,
	}

	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", settings.port),
		Handler:      appInstance.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "address", apiServer.Addr, "env", settings.env)
	err := apiServer.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
