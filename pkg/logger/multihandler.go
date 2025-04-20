package logger

import (
	"context"
	"log/slog"
)

type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(handlers ...slog.Handler) slog.Handler {
	return &MultiHandler{handlers: handlers}
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range m.handlers {
		// record.Time = record.Time.In(time.Local)
		_ = h.Handle(ctx, record)
	}
	return nil
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var hs []slog.Handler
	for _, h := range m.handlers {
		hs = append(hs, h.WithAttrs(attrs))
	}
	return NewMultiHandler(hs...)
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	var hs []slog.Handler
	for _, h := range m.handlers {
		hs = append(hs, h.WithGroup(name))
	}
	return NewMultiHandler(hs...)
}
