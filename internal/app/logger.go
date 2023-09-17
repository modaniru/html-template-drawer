package app

import (
	"log/slog"
	"os"
)

// TODO remove hardcode data
func configureLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
