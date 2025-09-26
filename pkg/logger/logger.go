package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

var Log *Extend

type PrettyHandlerOptions struct {
	SlogOpts   slog.HandlerOptions
	TimeFormat string
	UseColor   bool
	OutPutJson bool
}

type Extend struct {
	*slog.Logger
	handler *PrettyHandler
}

func (l *Extend) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *Extend) DebugMsgf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Debug(msg)
}

func (l *Extend) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Extend) InfoMsgf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Info(msg)
}

func (l *Extend) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}
func (l *Extend) WarnMsgf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Warn(msg)
}

func (l *Extend) Error(msg string, args ...any) error {
	l.Logger.Error(msg, args...)
	return fmt.Errorf(msg)
}

func (l *Extend) ErrorMsgf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Error(msg)
	return fmt.Errorf(msg)
}

type PrettyHandler struct {
	slog.Handler
	writer io.Writer
	opt    PrettyHandlerOptions
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	return &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		writer:  out,
		opt:     opts,
	}
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	// 时间部分
	timeStr := r.Time.Format(h.opt.TimeFormat)

	// 级别部分（带颜色）
	levelStr := r.Level.String() + ":"

	// 根据级别选择颜色函数
	var coloredLevel string
	if h.opt.UseColor {
		switch r.Level {
		case slog.LevelDebug:
			coloredLevel = color.MagentaString(levelStr)
		case slog.LevelInfo:
			coloredLevel = color.BlueString(levelStr)
		case slog.LevelWarn:
			coloredLevel = color.YellowString(levelStr)
		case slog.LevelError:
			coloredLevel = color.RedString(levelStr)
		default:
			coloredLevel = levelStr
		}
	} else {
		coloredLevel = levelStr
	}

	// 构建基础行
	output := fmt.Sprintf("%s %s %s",
		timeStr,
		coloredLevel,
		r.Message)

	// 处理附加字段
	if r.NumAttrs() > 0 {
		fields := make(map[string]interface{})
		r.Attrs(func(attr slog.Attr) bool {
			fields[attr.Key] = attr.Value.Any()
			return true
		})

		if h.opt.OutPutJson {
			if jsonData, err := json.MarshalIndent(fields, "", "  "); err == nil {
				output += "\n" + string(jsonData)
			}
		} else {
			// 简单键值对格式
			for k, v := range fields {
				output += fmt.Sprintf(" %s=%v", k, v)
			}
		}
	}

	// 确保换行
	output += "\n"

	// 直接输出到writer
	_, err := h.writer.Write([]byte(output))
	return err
}

func Init(option ...Option) {
	opts := PrettyHandlerOptions{
		SlogOpts:   slog.HandlerOptions{Level: slog.LevelDebug},
		TimeFormat: "2006-01-02 15:04:05",
		UseColor:   true,
		OutPutJson: false,
	}

	for _, opt := range option {
		opt(&opts)
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	Log = &Extend{
		Logger:  slog.New(handler),
		handler: handler,
	}
}
