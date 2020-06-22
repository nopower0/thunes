package settings

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogger() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/thunes.log",
		MaxSize:    1024, // in MB
		MaxBackups: 5,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	zap.ReplaceGlobals(logger)
}
