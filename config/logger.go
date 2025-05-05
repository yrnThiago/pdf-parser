package config

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

const logsPath = "./internal/logs/logs.log"

type LogLevels struct {
	Debug string `json:"DEBUG"`
}

func LoggerInit() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)

	if !pathExists(logsPath) {
		createFileWithPath(logsPath)
	}

	logFile, _ := os.OpenFile(logsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileWriter := zapcore.AddSync(logFile)
	consoleWriter := zapcore.AddSync(os.Stdout)

	logLevel := getLogLevel(Env.LogLevel)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, logLevel),
		zapcore.NewCore(consoleEncoder, consoleWriter, logLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func getLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "PROD":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

func createFileWithPath(filePath string) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}

	dir := filepath.Dir(absPath)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create parent folders: %v", err)
	}

	file, err := os.Create(absPath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()
}
