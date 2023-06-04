package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger represents a logger using zerolog.
type Logger struct {
	logger zerolog.Logger
}

// NewLogger creates a new Logger instance.
func NewLogger() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	output.FormatCaller = func(i interface{}) string {
		return ""
	}
	logger := zerolog.New(output).With().Timestamp().Logger()

	return &Logger{logger: logger}
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
	if os.Getenv("TEST_MODE") != "true" {
		l.logger.Info().Msg(message)
	}
}

// Error logs an error message.
func (l *Logger) Error(message string) {
	l.logger.Error().Msg(message)
}
