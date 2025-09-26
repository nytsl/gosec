package logger

import (
	"log/slog"
)

type Option func(*PrettyHandlerOptions)

func WithLevel(level slog.Level) Option {
	return func(o *PrettyHandlerOptions) {
		o.SlogOpts.Level = level
	}
}

func WithTimeFormat(format string) Option {
	return func(o *PrettyHandlerOptions) {
		o.TimeFormat = format // 正确设置时间格式
	}
}

func WithOutputJson(outputJson bool) Option {
	return func(o *PrettyHandlerOptions) {
		o.OutPutJson = outputJson
	}
}

func WithUseColor(useColor bool) Option {
	return func(o *PrettyHandlerOptions) {
		o.UseColor = useColor
	}
}
