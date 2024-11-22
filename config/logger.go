package config

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger exposes the zap logger.
var Logger = initLogger()

func initLogger() *zap.Logger {
	logConfig := zap.NewProductionConfig()

	logConfig.Encoding = "json"

	logConfig.EncoderConfig = zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		MessageKey:   "message",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder, // "level":"ERROR"
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	l, err := logConfig.Build()
	if err != nil {
		log.Fatalf("initLogger: %v", err)
	}

	return l
}

// Print prints to the logger.
func Print(v ...any) {
	Logger.Info(fmt.Sprint(v...))
}

// Fatal is equivalent to Print() followed by os.Exit(1).
func Fatal(v ...any) {
	Logger.Fatal(fmt.Sprint(v...))
}

// Printf prints to the logger.
func Printf(format string, v ...any) {
	Logger.Info(fmt.Sprintf(format, v...))
}

// Fatalf is equivalent to Printf() followed by os.Exit(1).
func Fatalf(format string, v ...any) {
	Logger.Fatal(fmt.Sprintf(format, v...))
}

// Errorf prints error to the logger.
func Errorf(format string, v ...any) {
	Logger.Error(fmt.Sprintf(format, v...))
}
