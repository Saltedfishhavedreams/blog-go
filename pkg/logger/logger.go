package logger

import (
	"blog/config"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(logConfig *config.LogConfig, mode string) (err error) {
	encoderConfig := getZapLogEncoder()
	writeSync := getFileLogWriter(logConfig)

	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte(logConfig.Level))

	if err != nil {
		return
	}

	var core zapcore.Core
	if mode == "dev" {
		// consoleEncoder := zapcore.NewConsoleEncoder(encoder)
		jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
		core = zapcore.NewTee(
			zapcore.NewCore(jsonEncoder, writeSync, level),
			// 终端日志输出
			zapcore.NewCore(jsonEncoder, zapcore.Lock(os.Stdout), level),
		)
	} else {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSync, level)
	}

	logger = zap.New(core)
	return
}

// 定制 encoderConfig
func getZapLogEncoder() zapcore.EncoderConfig {
	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "time"
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	return encoder
}

// 日志切割
func getFileLogWriter(logConfig *config.LogConfig) (writeSync zapcore.WriteSyncer) {
	fileWriter := &lumberjack.Logger{
		Filename:   logConfig.Filename,
		MaxSize:    logConfig.MaxSize,
		MaxAge:     logConfig.MaxAge,
		MaxBackups: logConfig.MaxBackups,
		Compress:   logConfig.Compress,
	}

	writeSync = zapcore.AddSync(fileWriter)
	return
}
