package logging

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger() *Logger {
	handler := slog.NewTextHandler(os.Stdout, nil)

	return &Logger{
		logger: slog.New(handler),
	}
}

func (l *Logger) Info(message string, args ...any) {
	l.logger.Info(message, args...)
}

func (l *Logger) Warn(message string, args ...any) {
	l.logger.Warn(message, args...)
}

func (l *Logger) Error(message string, args ...any) {
	l.logger.Error(message, args...)
}
