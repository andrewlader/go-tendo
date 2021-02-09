package main

import (
	"flag"
	"log"
	"strings"

	"github.com/andrewlader/go-tendo/tendo"
)

var path string
var languageType tendo.LanguageType
var logLevel tendo.LogLevel

func init() {
	parseArguments()
}

func main() {
	tendo := tendo.NewTendo(logLevel)
	tendo.Inspect(path, languageType)

	tendo.DisplayTotals()
}

func parseArguments() {
	var logLevelFlag string
	var logLevelmapping = map[string]tendo.LogLevel{
		"all":      tendo.LogAll,
		"trace":    tendo.LogTrace,
		"info":     tendo.LogInfo,
		"warnings": tendo.LogWarnings,
		"errors":   tendo.LogErrors,
	}

	if len(flag.Args()) < 1 {
		log.Fatal("Failed to provide enough arguments: no path was provided")
	}

	flag.StringVar(&logLevelFlag, "log", "all", "defines the level for logging output")
	flag.Parse()

	path = flag.Arg(0)
	languageType = tendo.LanguageType(tendo.Golang)
	logLevel = logLevelmapping[strings.ToLower(logLevelFlag)]
}
