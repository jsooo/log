package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type (
	Level int
)

const (
	LevelFatal = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

var (
	_fatal   *logger
	_error   *logger
	_warning *logger
	_info    *logger
	_debug   *logger
)

func init() {
	_fatal = &logger{_log: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile), logLevel: LevelFatal}
	_error = &logger{_log: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile), logLevel: LevelError}
	_warning = &logger{_log: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile), logLevel: LevelWarning}
	_info = &logger{_log: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile), logLevel: LevelInfo}
	_debug = &logger{_log: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile), logLevel: LevelDebug}
}

func Fatal(s string) {
	_fatal.Output(LevelFatal, s)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	_fatal.Output(LevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(s string) {
	_error.Output(LevelError, s)
}

func Errorf(format string, v ...interface{}) {
	_error.Output(LevelError, fmt.Sprintf(format, v...))
}

func Warn(s string) {
	_warning.Output(LevelWarning, s)
}

func Warnf(format string, v ...interface{}) {
	_warning.Output(LevelWarning, fmt.Sprintf(format, v...))
}

func Info(s string) {
	_info.Output(LevelInfo, s)
}

func Infof(format string, v ...interface{}) {
	_info.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func Debug(s string) {
	_debug.Output(LevelDebug, s)
}

func Debugf(format string, v ...interface{}) {
	_debug.Output(LevelDebug, fmt.Sprintf(format, v...))
}

func SetLogOutput(w io.Writer, level Level) {
	logger := getLogger(level)
	logger.SetOutput(w)
}

func SetLogPrefix(prefix string, level Level) {
	logger := getLogger(level)
	logger.SetPrefix(prefix)
}

func SetLogFlag(flag int, level Level) {
	logger := getLogger(level)
	logger.SetFlags(flag)
}

func getLogger(level Level) *logger {
	switch level {
	case LevelFatal:
		return _fatal
	case LevelError:
		return _error
	case LevelWarning:
		return _warning
	case LevelInfo:
		return _info
	case LevelDebug:
		return _fatal
	}

	return nil
}

type logger struct {
	_log *log.Logger
	//小于等于该级别的level才会被记录
	logLevel Level
}

func (l *logger) Output(level Level, s string) error {
	if l.logLevel < level {
		return nil
	}
	formatStr := "[UNKNOWN] %s"
	switch level {
	case LevelFatal:
		formatStr = "\033[35m[FATAL]\033[0m %s"
	case LevelError:
		formatStr = "\033[31m[ERROR]\033[0m %s"
	case LevelWarning:
		formatStr = "\033[33m[WARN]\033[0m %s"
	case LevelInfo:
		formatStr = "\033[32m[INFO]\033[0m %s"
	case LevelDebug:
		formatStr = "\033[36m[DEBUG]\033[0m %s"
	}
	s = fmt.Sprintf(formatStr, s)
	return l._log.Output(3, s)
}

func (l *logger) SetOutput(w io.Writer) {
	l._log.SetOutput(w)
}

func (l *logger) SetFlags(flag int) {
	l._log.SetFlags(flag)
}

func (l *logger) SetPrefix(prefix string) {
	l._log.SetPrefix(prefix)
}
