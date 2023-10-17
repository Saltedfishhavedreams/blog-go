package logger

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"
)

// 定制 Caller
func getCallerInfoConfig() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2)

	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	callerFields = append(callerFields, zap.String("file", fmt.Sprintf("%s:%d", file, line)), zap.String("func", funcName))

	return
}

func Info(msg string, fields ...zap.Field) {
	callerFields := getCallerInfoConfig()
	fields = append(fields, callerFields...)
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	callerFields := getCallerInfoConfig()
	fields = append(fields, callerFields...)
	logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	callerFields := getCallerInfoConfig()
	fields = append(fields, callerFields...)
	logger.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	callerFields := getCallerInfoConfig()
	fields = append(fields, callerFields...)
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	callerFields := getCallerInfoConfig()
	fields = append(fields, callerFields...)
	logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	callerFields := getCallerInfoConfig()
	fields = append(fields, callerFields...)
	logger.Panic(msg, fields...)
}
