package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
	initOnce      sync.Once
)

// Initialise logger once when logger is called.
func init() {
	initOnce.Do(func() {
		configureLogger()
	})
}

// Configure logger to write to file if env variable is set.
func configureLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	if len(env.LOG_FILE) > 0 {
		setupFileLogger(&config, env.LOG_FILE)
	} else {
		setupDefaultLogger(&config)
	}
}

// Setup logger with file output.
func setupFileLogger(config *zap.Config, logFilePath string) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	fullPath := filepath.Join(logDir, logFilePath)
	file, err := os.Create(fullPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to create log file: %v", err))
	}

	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(file))
	encoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
	core := zapcore.NewCore(encoder, writeSyncer, config.Level)
	logger = zap.New(core)
	sugaredLogger = logger.Sugar()
}

// Setup default logger without file output.
func setupDefaultLogger(config *zap.Config) {
	var err error
	logger, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialise zap logger: %v", err))
	}
	sugaredLogger = logger.Sugar()
}

func Info(msg string) {
	sugaredLogger.Info(msg)
}

func Infof(msg string, args ...interface{}) {
	sugaredLogger.Infof(msg, args...)
}

func Error(msg string) {
	sugaredLogger.Error(msg)
}

func Errorf(msg string, args ...interface{}) {
	sugaredLogger.Errorf(msg, args...)
}

func Warning(msg string) {
	sugaredLogger.Warn(msg)
}

func Warningf(msg string, args ...interface{}) {
	sugaredLogger.Warnf(msg, args...)
}

func Debug(msg string) {
	sugaredLogger.Debug(msg)
}

func Debugf(msg string, args ...interface{}) {
	sugaredLogger.Debugf(msg, args...)
}

func Panic(msg string) {
	sugaredLogger.Panic(msg)
}

func Panicf(msg string, args ...interface{}) {
	sugaredLogger.Panicf(msg, args...)
}
