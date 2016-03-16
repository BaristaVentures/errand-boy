package logging

import "github.com/Sirupsen/logrus"

// Logger identifies a loggable struct.
type Logger interface {
	GetContext() logrus.Fields
}

// Info logs a message at info level.
func Info(logger Logger, format string, args ...interface{}) {
	logrus.WithFields(logger.GetContext()).Infof(format, args)
}

// Error logs a message at error level.
func Error(logger Logger, format string, args ...interface{}) {
	logrus.WithFields(logger.GetContext()).Errorf(format, args)
}
