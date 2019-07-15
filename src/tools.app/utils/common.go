package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// GetCurrentDateTime returns current timestamp as "YY_MM_DD_HH_MM".
func GetCurrentDateTime() string {
	t := time.Now()
	return fmt.Sprintf("%d_%d_%d_%d_%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
}

// NewLoggerToStdout returns a custom logger write to stdout.
func NewLoggerToStdout(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}

// NewLoggerToFile returns a custom logger write to specified file.
func NewLoggerToFile(path, prefix string) (*log.Logger, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return log.New(f, prefix, log.Ldate|log.Ltime|log.Lshortfile), nil
}
