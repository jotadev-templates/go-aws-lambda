package bootstrap

import (
	"log/slog"
	"os"
)

func newLogger() *slog.Logger {
	const (
		envLogLevel string = "APP_LOG_LEVEL"
		debugLevel  string = "DEBUG"
	)

	loggerOpts := &slog.HandlerOptions{Level: slog.LevelInfo}

	level, ok := os.LookupEnv(envLogLevel)
	if ok && level == debugLevel {
		loggerOpts.Level = slog.LevelDebug
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, loggerOpts))
}
