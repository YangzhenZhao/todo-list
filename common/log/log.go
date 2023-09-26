package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger

func NewLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}
	newLogger := logrus.New()
	newLogger.SetOutput(&lumberjack.Logger{
		Filename:  "/home/ubuntu/todo-list/logs/todo.log",
		MaxSize:   1, // MB
		LocalTime: true,
		MaxAge:    28, //days
	})

	logger = newLogger
	if logger == nil {
		panic("Init logger failed!!!")
	}
	return logger
}
