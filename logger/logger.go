package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"time"
)

type Config struct {
	Output  string
	Level   string
	Service string
	Version string
}

// Logger defines the interface for logging operations.
type Logger interface {
	Info(ctx context.Context, msg string, data interface{})
	Warn(ctx context.Context, msg string, data interface{})
	Error(ctx context.Context, msg string, err error)
	Debug(ctx context.Context, msg string, data interface{})
}

// logger implements the Logger interface.
type logger struct {
	logger *slog.Logger
}

// Info logs an informational message with optional fields.
func (l *logger) Info(ctx context.Context, msg string, data interface{}) {
	l.logger.With(slog.String("type", "sys")).InfoContext(ctx, msg, slog.Any("data", data))
}

// Warn logs a warning message with optional fields.
func (l *logger) Warn(ctx context.Context, msg string, data interface{}) {
	l.logger.With(slog.String("type", "sys")).WarnContext(ctx, msg, slog.Any("data", data))
}

// Error logs an error message with optional fields.
func (l *logger) Error(ctx context.Context, msg string, err error) {
	l.logger.With(slog.String("type", "sys")).ErrorContext(ctx, msg, slog.String("error", err.Error()))
}

// Debug logs a debug message with optional fields.
func (l *logger) Debug(ctx context.Context, msg string, data interface{}) {
	l.logger.With(slog.String("type", "sys")).DebugContext(ctx, msg, slog.Any("data", data))
}

// New creates a new Logger instance with default settings.
func New(cfg Config) Logger {
	var handler slog.Handler

	opt := slog.HandlerOptions{
		AddSource: false,
		Level:     parseLevel(cfg.Level),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				a.Value = slog.Int64Value(time.Now().Unix())
			}
			return a
		},
	}

	switch strings.ToLower(cfg.Output) {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &opt)
	default:
		handler = slog.NewTextHandler(os.Stdout, &opt)
	}

	l := slog.New(handler).
		With(slog.String("service", cfg.Service)).
		With(slog.String("version", cfg.Version))
	slog.SetDefault(l)

	return &logger{
		logger: l,
	}
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "debug":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
