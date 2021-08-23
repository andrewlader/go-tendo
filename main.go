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
	var languageTypeFlag string
	var logLevelmapping = map[string]tendo.LogLevel{
		"all":      tendo.LogAll,
		"trace":    tendo.LogTrace,
		"info":     tendo.LogInfo,
		"warnings": tendo.LogWarnings,
		"errors":   tendo.LogErrors,
	}
	var languageTypemapping = map[string]tendo.LanguageType{
		"go":   tendo.LanguageType(tendo.Golang),
		"java": tendo.LanguageType(tendo.Java),
	}

	flag.StringVar(&path, "path", "./", "defines the path to walk")
	flag.StringVar(&logLevelFlag, "log", "all", "defines the level for logging output")
	flag.StringVar(&languageTypeFlag, "language", "go", "defines the programming language for the path")

	flag.Parse()

	languageType = tendo.LanguageType(languageTypemapping[languageTypeFlag])
	logLevel = logLevelmapping[strings.ToLower(logLevelFlag)]

	log.Printf("path: %s, language type: %s, log level: %s", path, languageTypeFlag, logLevelFlag)
}
