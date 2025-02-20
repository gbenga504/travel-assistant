package utils

import (
	"log/slog"
	"os"
)

type LoggerParams struct {
	msg      string
	category string
	context  struct {
		userId *string
	}
	payload map[string]any
}

type ErrorParams struct {
	name    string
	message string
}

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}

func Info(params LoggerParams) {
	logger := NewLogger()

	logger.Info(params.msg,
		slog.String("category", params.category),
		slog.Group("labels", slog.String("application", "api")),
		slog.Group("context", slog.String("userId", *params.context.userId)),
		slog.Group("payload", sliceToAnySlice(argsToSlogAttrSlice(params.payload))...))
}

func Warn(params LoggerParams) {
	logger := NewLogger()

	logger.Warn(params.msg,
		slog.String("category", params.category),
		slog.Group("labels", slog.String("application", "api")),
		slog.Group("context", slog.String("userId", *params.context.userId)),
		slog.Group("payload", sliceToAnySlice(argsToSlogAttrSlice(params.payload))...))
}

func Error(err ErrorParams, params LoggerParams) {
	logger := NewLogger()

	logger.Error(params.msg,
		slog.Group("error", slog.String("name", err.name), slog.String("message", err.message)),
		slog.String("category", params.category),
		slog.Group("labels", slog.String("application", "api")),
		slog.Group("context", slog.String("userId", *params.context.userId)),
		slog.Group("payload", sliceToAnySlice(argsToSlogAttrSlice(params.payload))...))
}

func argsToSlogAttrSlice(args map[string]any) []slog.Attr {
	result := make([]slog.Attr, len(args))

	for key, value := range args {
		result = append(result, slog.Any(key, value))
	}

	return result
}

func sliceToAnySlice(slogAttrs []slog.Attr) []any {
	result := make([]any, len(slogAttrs))

	for _, value := range slogAttrs {
		result = append(result, value)
	}

	return result
}
