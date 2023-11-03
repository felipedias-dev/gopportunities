package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	debug *log.Logger
	info *log.Logger
	warning *log.Logger
	err *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	debugPrefix := fmt.Sprintf("[%s] - DEBUG: ", prefix)
	infoPrefix := fmt.Sprintf("[%s] - INFO: ", prefix)
	warningPrefix := fmt.Sprintf("[%s] - WARNING: ", prefix)
	errorPrefix := fmt.Sprintf("[%s] - ERROR: ", prefix)

	return &Logger{
		debug: log.New(writer, debugPrefix, logger.Flags()),
		info: log.New(writer, infoPrefix, logger.Flags()),
		warning: log.New(writer, warningPrefix, logger.Flags()),
		err: log.New(writer, errorPrefix, logger.Flags()),
		writer: writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
