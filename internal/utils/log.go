package utils

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger *log.Logger
	WarningLogger *log.Logger
	errorLogger *log.Logger
	criticalLogger *log.Logger
)

func InitLogger() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	criticalLogger = log.New(os.Stdout, "[CRITICAL] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(context, msg string, args ...any) {
	infoLogger.Printf("[%s] %s", context, format(msg, args...))
}

func LogWarning(context, msg string, args ...any) {
	WarningLogger.Printf("[%s] %s", context, format(msg, args...))
}

func LogError(context, msg string, args ...any) {
	errorLogger.Printf("[%s] %s", context, format(msg, args...))
}

func LogCritical(context, msg string, args ...any) {
	criticalLogger.Printf("[%s] %s", context, format(msg, args...))
}

func format(msg string, args ...any) string {
	if len(args) > 0 {
		return sprintf(msg, args...)
	}
	return msg
}

func sprintf(msg string, args ...any) string {
	return fmt.Sprintf(msg, args...)
}
