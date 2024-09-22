package logger

import (
	"log/slog"

	"github.com/profclems/go-dotenv"
)

func LogLevel(level string) slog.Level {
	envLevel := dotenv.GetString("LOG_LEVEL")

	switch envLevel {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	case "FATAL":
		return slog.LevelError
	default:
		slog.Info("Log level set to DEFAULT=DEBUG")
		return slog.LevelDebug
	}
}
