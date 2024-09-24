package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Set() {
	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	slog.SetDefault(logger)
}
