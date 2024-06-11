package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.LstdFlags)

	return &Logger{
		debug:   log.New(writer, fmt.Sprintf("%s %s %s: ", blue, "DEBUG", reset), logger.Flags()),
		info:    log.New(writer, fmt.Sprintf("%s %s %s: ", green, "INFO", reset), logger.Flags()),
		warning: log.New(writer, fmt.Sprintf("%s %s %s: ", yellow, "WARNING", reset), logger.Flags()),
		error:   log.New(writer, fmt.Sprintf("%s %s %s: ", red, "ERROR", reset), logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Print(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Print(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Print(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Print(v...)
}
