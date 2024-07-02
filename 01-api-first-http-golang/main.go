package main

import (
	"context"
	"log/slog"
	"os"

	"apifirst/app"
	"apifirst/ports/httpapi"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	application := app.NewApp()

	srv, err := httpapi.NewServer(context.Background(), 8080, application, logger)
	if err != nil {
		logger.Error("error creating/configuring the http server", "error", err)
		return
	}

	err = srv.Start()
	if err != nil {
		logger.Error("the http server crashed", "error", err)
		os.Exit(1)
	}

	// TODO: graceful termination
}
