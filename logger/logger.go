package logger

import (
	"io"
	"log"
	"os"
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

	flag := log.LstdFlags | log.Lmsgprefix | log.Lshortfile

	return &logger{
		errorLogger:   log.New(out, "ERROR: ", flag),
		infoLogger:    log.New(out, "INFO: ", flag),
		debugLogger:   log.New(out, "DEBUG: ", flag),
		traceLogger:   log.New(out, "TRACE: ", flag),
		fatalLogger:   log.New(out, "FATAL: ", flag),
		panicLogger:   log.New(out, "PANIC: ", flag),
		defaultLogger: log.New(out, "", flag),
	}
}

func New() Logger {
	return NewWithWriter(os.Stdout)
}

func (l *logger) Error(v ...any) {
	l.errorLogger.Println(v...)
}

func (l *logger) Info(v ...any) {
	l.infoLogger.Println(v...)
}

func (l *logger) Debug(v ...any) {
	l.debugLogger.Println(v...)
}

func (l *logger) Trace(v ...any) {
	l.traceLogger.Println(v...)
}

func (l *logger) Fatal(v ...any) {
	l.fatalLogger.Fatalln(v...)
}

func (l *logger) Panic(v ...any) {
	l.panicLogger.Panicln(v...)
}

func (l *logger) Print(v ...any) {
	l.defaultLogger.Println(v...)
}

func (l *logger) Printf(format string, v ...any) {
	l.defaultLogger.Printf(format, v...)
}

func (l *logger) Println(v ...any) {
	l.defaultLogger.Println(v...)
}
