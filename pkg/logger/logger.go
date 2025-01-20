package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger initializes a new logger with specified output and file options.
func NewLogger(logFile string) (*Logger, error) {
	var file io.Writer = os.Stdout

	if logFile != "" {
		var err error
		file, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
	}

	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}, nil
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

// Error logs an error message.
func (l *Logger) Error(err error) {
	l.errorLogger.Println(err)
}
