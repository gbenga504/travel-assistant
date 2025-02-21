package logger

import "log/slog"

type Context struct {
	UserId string
}

type Options struct {
	Category string
	Context  Context
	Payload  map[string]any
}

type Option func(*Options)

func WithCategory(category string) Option {
	return func(opt *Options) {
		opt.Category = category
	}
}

func WithContext(context Context) Option {
	return func(opt *Options) {
		opt.Context = context
	}
}

func WithPayload(payload map[string]any) Option {
	return func(opt *Options) {
		opt.Payload = payload
	}
}

func getOptions(opts ...Option) *Options {
	options := Options{Category: "core", Context: Context{}}

	for _, opt := range opts {
		opt(&options)
	}

	return &options
}

func contextToSlogAttrSlice(args Context) []slog.Attr {
	var result []slog.Attr

	if args.UserId != "" {
		result = append(result, slog.String("userId", args.UserId))
	}

	return result
}

func mapToSlogAttrSlice(args map[string]any) []slog.Attr {
	var result []slog.Attr

	for key, value := range args {
		result = append(result, slog.Any(key, value))
	}

	return result
}

func slogAttrSliceToAnySlice(slogAttrs []slog.Attr) []any {
	var result []any

	for _, value := range slogAttrs {
		result = append(result, value)
	}

	return result
}
