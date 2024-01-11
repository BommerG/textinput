package main

import (
	"log"
	"log/slog"
	"os"
)

var (
	logfile = "domino.log"
	logger  *slog.Logger
)

func init() {
	fh, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		log.Fatal(err)
	}

	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(fh, opts)
	logger = slog.New(handler)
	slog.SetDefault(logger)
}
