package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.EncoderConfig.CallerKey = zapcore.OmitKey
	config.DisableStacktrace = true
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, _ := config.Build()
	logger = zapLogger
}

func Debug(message string, fields ...zapcore.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zapcore.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zapcore.Field) {
	logger.Warn(message, fields...)
}

func Error(message string, err error, fields ...zapcore.Field) {
	fields = append(fields, zap.Error(err))
	logger.Error(message, fields...)
}

func Fatal(message string, err error, fields ...zapcore.Field) {
	fields = append(fields, zap.Error(err))
	logger.Fatal(message, fields...)
}
