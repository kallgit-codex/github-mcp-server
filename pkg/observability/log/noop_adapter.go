package log

import (
	"context"
	"log/slog"
)

type NoopLogger struct{}

var _ Logger = (*NoopLogger)(nil)

func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}

func (l *NoopLogger) Level() Level {
	return DebugLevel
}

func (l *NoopLogger) Log(ctx context.Context, level Level, msg string, fields ...slog.Attr) {
	// No-op
}

func (l *NoopLogger) Debug(msg string, fields ...slog.Attr) {
	// No-op
}

func (l *NoopLogger) Info(msg string, fields ...slog.Attr) {
	// No-op
}

func (l *NoopLogger) Warn(msg string, fields ...slog.Attr) {
	// No-op
}

func (l *NoopLogger) Error(msg string, fields ...slog.Attr) {
	// No-op
}

func (l *NoopLogger) Fatal(msg string, fields ...slog.Attr) {
	// No-op
}

func (l *NoopLogger) WithFields(fields ...slog.Attr) Logger {
	return l
}

func (l *NoopLogger) WithError(err error) Logger {
	return l
}

func (l *NoopLogger) Named(name string) Logger {
	return l
}

func (l *NoopLogger) WithLevel(level Level) Logger {
	return l
}

func (l *NoopLogger) Sync() error {
	return nil
}
