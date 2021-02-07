package tendo

import (
	"fmt"
	"log"
)

// LogLevel determines how much to output when logging
type LogLevel uint8

// The various levels of logging
const (
	LogAll LogLevel = iota
	LogTrace
	LogInfo
	LogWarnings
	LogErrors
	logAlways // internal use only
)

type Logger struct {
	level LogLevel
}

func newLogger(level LogLevel) *Logger {
	log.SetPrefix("{go-tendo} ")

	if (level < LogAll) || (level > logAlways) {
		return nil
	}

	return &Logger{
		level: level,
	}
}

func (theLogger *Logger) println(logLevel LogLevel, output string) {
	if logLevel >= theLogger.level {
		fmt.Println(output)
	}
}

func (theLogger *Logger) printfln(logLevel LogLevel, format string, args ...interface{}) {
	if logLevel >= theLogger.level {
		output := fmt.Sprintf(format, args...)
		fmt.Println(output)
	}
}

func (theLogger *Logger) printf(logLevel LogLevel, format string, args ...interface{}) {
	if logLevel >= theLogger.level {
		log.Print(fmt.Sprintf(format, args...))
	}
}

func (theLogger *Logger) fatalf(logLevel LogLevel, format string, args ...interface{}) {
	if logLevel >= theLogger.level {
		log.Fatal(fmt.Sprintf(format, args...))
	}
}
