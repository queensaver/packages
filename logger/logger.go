package logger

// This is a logging module that enforces structured logging.

import (
	"flag"
	"os"
	"time"

	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugar   *zap.SugaredLogger
	verbose = flag.Bool("verbose", false, "Turn on more verbose logging")
)

func init() {
	flag.Parse()
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.UTC().Format(time.RFC3339Nano))
	}
	level := zapcore.InfoLevel
	if *verbose || os.Getenv("VERBOSE") != "" {
		level = zapcore.DebugLevel
	}
	logger := zap.New(zapcore.NewCore(
		zaplogfmt.NewEncoder(config),
		os.Stdout,
		level,
	))

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
