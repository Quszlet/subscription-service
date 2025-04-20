package logger

import (
	"log/slog"
)

type SlogAdapter struct {
	logger *slog.Logger
}

func NewSlogAdapter(l *slog.Logger) *SlogAdapter {
	return &SlogAdapter{logger: l}
}

func (l *SlogAdapter) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogAdapter) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *SlogAdapter) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *SlogAdapter) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}
