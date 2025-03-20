package logger

import "go.uber.org/zap"

type DefaultLogger struct {
	zap *zap.Logger
}

var DefaultLoggerInstance *DefaultLogger = NewDefaultLogger()
var _ Logger = (*DefaultLogger)(nil)

func NewDefaultLogger() *DefaultLogger {
	zapLogger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	return &DefaultLogger{
		zap: zapLogger,
	}
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
