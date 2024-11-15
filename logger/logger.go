package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	errorLogger   *log.Logger
	infoLogger    *log.Logger
	debugLogger   *log.Logger
	traceLogger   *log.Logger
	fatalLogger   *log.Logger
	panicLogger   *log.Logger
	defaultLogger *log.Logger
}

func New() *Logger {
	return NewWithWriter(os.Stdout)
}

func NewWithWriter(out io.Writer) *Logger {

	flag := log.LstdFlags | log.Lmsgprefix | log.Lshortfile

	return &Logger{
		errorLogger:   log.New(out, "ERROR: ", flag),
		infoLogger:    log.New(out, "INFO: ", flag),
		debugLogger:   log.New(out, "DEBUG: ", flag),
		traceLogger:   log.New(out, "TRACE: ", flag),
		fatalLogger:   log.New(out, "FATAL: ", flag),
		panicLogger:   log.New(out, "PANIC: ", flag),
		defaultLogger: log.New(out, "", flag),
	}
}

func (l *Logger) Error(v ...any) {
	l.errorLogger.Println(v...)
}

func (l *Logger) Info(v ...any) {
	l.infoLogger.Println(v...)
}

func (l *Logger) Debug(v ...any) {
	l.debugLogger.Println(v...)
}

func (l *Logger) Trace(v ...any) {
	l.traceLogger.Println(v...)
}

func (l *Logger) Fatal(v ...any) {
	l.fatalLogger.Fatalln(v...)
}

func (l *Logger) Panic(v ...any) {
	l.panicLogger.Panicln(v...)
}

func (l *Logger) Print(v ...any) {
	l.defaultLogger.Println(v...)
}

func (l *Logger) Printf(format string, v ...any) {
	l.defaultLogger.Printf(format, v...)
}

func (l *Logger) Println(v ...any) {
	l.defaultLogger.Println(v...)
}
