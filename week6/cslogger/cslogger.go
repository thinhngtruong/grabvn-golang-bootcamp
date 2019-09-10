package cslogger

import (
	"github.com/sirupsen/logrus"
	"io"
)

// TODO:
// - formatted in json - OK
// - can write logs to file - OK
// - can write logs to another file if current log file is to large
// - able for developer to trace between services

type Level int32

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

// CSLogger exported APIs
type CSLogger interface {
	SetOutput(out io.Writer)
	SetLogLevel(level Level)
	Info(args ...interface{})
	InvalidArg(argumentName string)
	InvalidArgValue(argumentName string, argumentValue string)
	MissingArg(argumentName string)
	Message(message string)
}

// NewCSLogger create new instance of CSLogger
func NewCSLogger() CSLogger {
	var logger = &csLogger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}

	return logger
}

type csLogger struct {
	*logrus.Logger
}

// InvalidArg is a standard error message
func (l *csLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage, argumentName)
}

// InvalidArgValue is a standard error message
func (l *csLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *csLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage, argumentName)
}

func (l *csLogger) Message(message string) {
	l.Errorf(myMessage, message)
}

func (l *csLogger) SetLogLevel(level Level) {
	lv := logrus.ErrorLevel

	switch level {
	case ErrorLevel:
		lv = logrus.ErrorLevel
	case WarnLevel:
		lv = logrus.WarnLevel
	case InfoLevel:
		lv = logrus.InfoLevel
	case DebugLevel:
		lv = logrus.DebugLevel
	}

	l.SetLevel(lv)
}
