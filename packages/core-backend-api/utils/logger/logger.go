package logger

import (
	"log/slog"
	"os"
)

type Context struct {
	UserId string
}

type LoggerParams struct {
	Msg      string
	Category *string
	Context  *Context
	Payload  map[string]any
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

	context := defaultContext(params.Context)

	logger.Info(params.Msg,
		slog.String("category", defaultCategory(params.Category)),
		slog.Group("labels", slog.String("application", "api")),
		slog.Group("context", slog.String("userId", context.UserId)),
		slog.Group("payload", sliceToAnySlice(argsToSlogAttrSlice(params.Payload))...))
}

// func Warn(params LoggerParams) {
// 	logger := NewLogger()

// 	logger.Warn(params.Msg,
// 		slog.String("category", params.Category),
// 		slog.Group("labels", slog.String("application", "api")),
// 		slog.Group("context", slog.String("userId", *params.Context.UserId)),
// 		slog.Group("payload", sliceToAnySlice(argsToSlogAttrSlice(params.Payload))...))
// }

// func Error(err ErrorParams, params LoggerParams) {
// 	logger := NewLogger()

// 	logger.Error(params.Msg,
// 		slog.Group("error", slog.String("name", err.name), slog.String("message", err.message)),
// 		slog.String("Category", params.Category),
// 		slog.Group("labels", slog.String("application", "api")),
// 		slog.Group("context", slog.String("userId", *params.Context.UserId)),
// 		slog.Group("payload", sliceToAnySlice(argsToSlogAttrSlice(params.Payload))...))
// }

func defaultCategory(category *string) string {
	result := "core"

	if category != nil {
		result = *category
	}

	return result
}

func defaultContext(context *Context) Context {
	result := Context{UserId: "NOT_DEFINED"}

	if context != nil {
		result = *context
	}

	return result
}

func argsToSlogAttrSlice(args map[string]any) []slog.Attr {
	var result []slog.Attr

	for key, value := range args {
		result = append(result, slog.Any(key, value))
	}

	return result
}

func sliceToAnySlice(slogAttrs []slog.Attr) []any {
	var result []any

	for _, value := range slogAttrs {
		result = append(result, value)
	}

	return result
}
