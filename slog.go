package ctxslog

import (
	"context"
	"log/slog"
)

var getDefaultLogger = slog.Default

// DisableDefaultLogger disables the default logger.
// After calling this function, calls to FromContext will return nil
// if no logger is set in the context.
func DisableDefaultLogger() {
	getDefaultLogger = func() *slog.Logger {
		return nil
	}
}

// EnableDefaultLogger enables the default logger.
// After calling this function, calls to FromContext will return the default logger
// if no logger is set in the context.
func EnableDefaultLogger() {
	getDefaultLogger = slog.Default
}

type LoggerKey string

// NewContext creates a new context with the provided logger.
// If the context is nil, it creates a new background context.
// The logger is stored in the context under the key "logger".
func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, LoggerKey("logger"), logger)
}

// FromContext retrieves the logger from the context.
// If no logger is found, it returns the default logger if enabled,
// or nil if the default logger is disabled.
func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(LoggerKey("logger")).(*slog.Logger)
	if !ok {
		return getDefaultLogger()
	}
	return logger
}

// With returns a Logger that includes the given attributes
// in each output operation. Arguments are converted to
// attributes as if by [Logger.Log].
func With(ctx context.Context, args ...any) *slog.Logger {
	l := FromContext(ctx)
	if l == nil {
		return nil
	}
	return l.With(args...)
}

// WithGroup returns a Logger that starts a group, if name is non-empty.
// The keys of all attributes added to the Logger will be qualified by the given
// name. (How that qualification happens depends on the [Handler.WithGroup]
// method of the Logger's Handler.)
//
// If name is empty, WithGroup returns the receiver.
func WithGroup(ctx context.Context, name string) *slog.Logger {
	l := FromContext(ctx)
	if l == nil {
		return nil
	}
	return l.WithGroup(name)
}

// Debug calls [Logger.Debug] on the logger in the context.
func Debug(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Debug(msg, args...)
}

// DebugContext calls [Logger.DebugContext] on the logger in the context.
func DebugContext(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.DebugContext(ctx, msg, args...)
}

// Info calls [Logger.Info] on the logger in the context.
func Info(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Info(msg, args...)
}

// InfoContext calls [Logger.InfoContext] on the logger in the context.
func InfoContext(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.InfoContext(ctx, msg, args...)
}

// Warn calls [Logger.Warn] on the logger in the context.
func Warn(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Warn(msg, args...)
}

// WarnContext calls [Logger.WarnContext] on the logger in the context.
func WarnContext(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.WarnContext(ctx, msg, args...)
}

// Error calls [Logger.Error] on the logger in the context.
func Error(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Error(msg, args...)
}

// ErrorContext calls [Logger.ErrorContext] on the logger in the context.
func ErrorContext(ctx context.Context, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.ErrorContext(ctx, msg, args...)
}

// Log calls [Logger.Log] on the logger in the context.
func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.Log(ctx, level, msg, args...)
}

// LogAttrs calls [Logger.LogAttrs] on the logger in the context.
func LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l := FromContext(ctx)
	if l == nil {
		return
	}
	l.LogAttrs(ctx, level, msg, attrs...)
}
