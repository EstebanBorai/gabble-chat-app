package logger

import "log"

// LogLevel is a preset for logging levels
type LogLevel int8

const (
	// Info will log any message received
	Info LogLevel = 3
	// Warning will log warning messages and error messages
	Warning LogLevel = 2
	// Error will log errors only
	Error LogLevel = 1
)

// ANSI Colors
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

// Logger for Chat runtime with configurable log levels
type Logger struct {
	level LogLevel
}

// NewLogger creates a new Logger instance
func NewLogger(level int8) *Logger {
	logger := new(Logger)
	logLevel := LogLevel(level)

	logger.level = logLevel

	return logger
}

// Info logs "info" level messages
func (l *Logger) Info(message string) {
	l.log(Info, message)
}

// Warning logs "warning" level messages
func (l *Logger) Warning(message string) {
	l.log(Warning, message)
}

// Error logs "error" level messages
func (l *Logger) Error(message error) {
	str := message.Error()
	l.log(Error, str)
}

func (l *Logger) log(level LogLevel, message string) {
	if level <= l.level {
		switch level {
		case Info:
			log.Printf("%s %s\n", InfoColor, message)
		case Warning:
			log.Printf("%s %s\n", WarningColor, message)
		case Error:
			log.Printf("%s %s\n", ErrorColor, message)
		default:
			log.Printf("%s\n", message)
		}
	}
}
