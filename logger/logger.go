package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type Logger interface {
	Error(v ...any)
	Info(v ...any)
	Debug(v ...any)
	Trace(v ...any)
	Fatal(v ...any)
	Panic(v ...any)
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
}

type logger struct {
	errorLogger   *log.Logger
	infoLogger    *log.Logger
	debugLogger   *log.Logger
	traceLogger   *log.Logger
	fatalLogger   *log.Logger
	panicLogger   *log.Logger
	defaultLogger *log.Logger
}

func NewWithWriter(out io.Writer) Logger {

	flag := log.Lmsgprefix

	color := func(c, prefix string) string {
		if os.Stdout == out {
			return fmt.Sprintf("%s%s: %s", c, prefix, colorReset)
		}
		return fmt.Sprintf("%s: ", prefix)
	}

	return &logger{
		errorLogger:   log.New(out, color(colorRed, "ERROR"), flag),
		infoLogger:    log.New(out, color(colorGreen, "INFO"), flag),
		debugLogger:   log.New(out, color(colorYellow, "DEBUG"), flag),
		traceLogger:   log.New(out, color(colorBlue, "TRACE"), flag),
		fatalLogger:   log.New(out, color(colorPurple, "FATAL"), flag),
		panicLogger:   log.New(out, color(colorCyan, "PANIC"), flag),
		defaultLogger: log.New(out, "", flag),
	}
}

func New() Logger {
	return NewWithWriter(os.Stdout)
}

func (l *logger) logWithCaller(logger *log.Logger, v ...any) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		prefix := fmt.Sprintf("%s:%d", file, line)
		logger.Println(append([]any{prefix}, v...)...)
	} else {
		logger.Println(v...)
	}
}

func (l *logger) Error(v ...any) {
	l.logWithCaller(l.errorLogger, v...)
}

func (l *logger) Info(v ...any) {
	l.logWithCaller(l.infoLogger, v...)
}

func (l *logger) Debug(v ...any) {
	l.logWithCaller(l.debugLogger, v...)
}

func (l *logger) Trace(v ...any) {
	l.logWithCaller(l.traceLogger, v...)
}

func (l *logger) Fatal(v ...any) {
	l.logWithCaller(l.fatalLogger, v...)
	os.Exit(1)
}

func (l *logger) Panic(v ...any) {
	l.logWithCaller(l.panicLogger, v...)
	panic(fmt.Sprint(v...))
}

func (l *logger) Print(v ...any) {
	l.logWithCaller(l.defaultLogger, v...)
}

func (l *logger) Printf(format string, v ...any) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		prefix := fmt.Sprintf("%s:%d: ", file, line)
		l.defaultLogger.Printf(prefix+format, v...)
	} else {
		l.defaultLogger.Printf(format, v...)
	}
}

func (l *logger) Println(v ...any) {
	l.logWithCaller(l.defaultLogger, v...)
}
