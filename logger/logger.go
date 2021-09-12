package logger

// This is a logging module that enforces structured logging.

import (
	"log"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar = logger.Sugar()
}

// TODO: Context logging: https://notes.burke.libbey.me/context-and-logging/

func Debug(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}

func Sync() {
	sugar.Sync()
}
