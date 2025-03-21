package logger

import (
	"sync"

	"go.uber.org/zap"
)

type DefaultLogger struct {
	zap *zap.Logger
}

var (
	defaultLoggerInstance *DefaultLogger
	once                  sync.Once
)

var _ Logger = (*DefaultLogger)(nil)

func GetDefaultLogger() *DefaultLogger {
	once.Do(func() {
		zapLogger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		defaultLoggerInstance = &DefaultLogger{
			zap: zapLogger,
		}
	})
	return defaultLoggerInstance
}

func (l *DefaultLogger) Debug(msg string) {
	l.zap.Debug(msg)
}

func (l *DefaultLogger) Info(msg string) {
	l.zap.Info(msg)
}

func (l *DefaultLogger) Warning(msg string) {
	l.zap.Warn(msg)
}

func (l *DefaultLogger) Error(msg string) {
	l.zap.Error(msg)
}

func (l *DefaultLogger) SetLevel(level int) {
	// TODO: Implement this
}
