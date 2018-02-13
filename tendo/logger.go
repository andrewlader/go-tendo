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

type logger struct {
	level LogLevel
}

func newLogger(level LogLevel) *logger {
	log.SetPrefix("{go-tendo} ")

	return &logger{
		level: level,
	}
}

func (logger *logger) println(logLevel LogLevel, output string) {
	if logLevel >= logger.level {
		fmt.Println(output)
	}
}

func (logger *logger) printfln(logLevel LogLevel, format string, args ...interface{}) {
	if logLevel >= logger.level {
		output := fmt.Sprintf(format, args...)
		fmt.Println(output)
	}
}

func (logger *logger) printf(logLevel LogLevel, format string, args ...interface{}) {
	if logLevel >= logger.level {
		log.Print(fmt.Sprintf(format, args...))
	}
}

func (logger *logger) fatalf(logLevel LogLevel, format string, args ...interface{}) {
	if logLevel >= logger.level {
		log.Fatal(fmt.Sprintf(format, args...))
	}
}
