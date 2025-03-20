package logger

type Logger interface {
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)

	SetLevel(int)
}
