package logger

type Ilogger interface {
	Debug(msg string, args ...any)
	DebugMsgf(format string, args ...interface{})
	Info(msg string, args ...any)
	InfoMsgf(format string, args ...interface{})
	Warn(msg string, args ...any)
	WarnMsgf(format string, args ...interface{})
	Error(msg string, args ...any) error
	ErrorMsgf(format string, args ...interface{}) error
}
