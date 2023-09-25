package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

func InitLogger() {
	logger := logrus.New()
	logger.SetOutput(&lumberjack.Logger{
		Filename: "../../logs/todo.log",
		MaxSize:  2,  // MB
		MaxAge:   28, //days
	})
	Logger = logger
}
