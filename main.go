package main

import (
	"flag"
	"log"
	"strings"

	"github.com/andrewlader/go-tendo/tendo"
)

func main() {
	path, logLevel := parseArguments()

	tendo := tendo.NewTendo(tendo.LanguageType(tendo.Golang), logLevel)
	tendo.Inspect(path)

	tendo.DisplayTotals()
}

func parseArguments() (string, tendo.LogLevel) {
	var logLevel string
	var logLevelmapping = map[string]tendo.LogLevel{
		"all":      tendo.LogAll,
		"trace":    tendo.LogTrace,
		"info":     tendo.LogInfo,
		"warnings": tendo.LogWarnings,
		"errors":   tendo.LogErrors,
	}

	flag.StringVar(&logLevel, "log", "all", "defines the level for logging output")
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Failed to provide enough arguments: no path was provided")
	}

	path := flag.Arg(0)
	loggingLevel := logLevelmapping[strings.ToLower(logLevel)]

	return path, loggingLevel
}
