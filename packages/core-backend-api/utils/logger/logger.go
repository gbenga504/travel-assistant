package logger

import (
	"log/slog"
	"os"
)

type ErrorOpt struct {
	Name    string
	Message string
}

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}

func Info(msg string, opts ...Option) {
	logger := NewLogger()
	options := getOptions(opts...)

	logger.Info(msg,
		slog.Group("labels", slog.String("application", "api"), slog.String("category", options.Category)),
		slog.Group("context", slogAttrSliceToAnySlice(contextToSlogAttrSlice(options.Context))...),
		slog.Group("payload", slogAttrSliceToAnySlice(mapToSlogAttrSlice(options.Payload))...))
}

func Warn(msg string, opts ...Option) {
	logger := NewLogger()
	options := getOptions(opts...)

	logger.Warn(msg,
		slog.Group("labels", slog.String("application", "api"), slog.String("category", options.Category)),
		slog.Group("context", slogAttrSliceToAnySlice(contextToSlogAttrSlice(options.Context))...),
		slog.Group("payload", slogAttrSliceToAnySlice(mapToSlogAttrSlice(options.Payload))...))
}

func Error(msg string, err ErrorOpt, opts ...Option) {
	logger := NewLogger()
	options := getOptions(opts...)

	logger.Warn(msg,
		slog.Group("labels", slog.String("application", "api"), slog.String("category", options.Category)),
		slog.Group("error", slog.String("name", err.Name), slog.String("message", err.Message)),
		slog.Group("context", slogAttrSliceToAnySlice(contextToSlogAttrSlice(options.Context))...),
		slog.Group("payload", slogAttrSliceToAnySlice(mapToSlogAttrSlice(options.Payload))...))
}
