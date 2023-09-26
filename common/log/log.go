package log

import (
	"log"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

func InitLogger() {
	logger := logrus.New()
	logger.SetOutput(&lumberjack.Logger{
		Filename:  "/Users/nocilantro/todo-list/logs/todo.log",
		MaxSize:   1, // MB
		LocalTime: true,
		MaxAge:    28, //days
	})

	Logger = logger
	if Logger == nil {
		log.Printf("invalid logger!!!")
	} else {
		Logger.Info("Init Logger success")
	}
}
