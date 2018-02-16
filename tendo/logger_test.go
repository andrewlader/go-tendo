package tendo

import "testing"

func TestLoggerPrintlnNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	logger := newLogger(LogInfo)
	logger.printfln(LogInfo, "2 + 2 = 4")
}

func TestLoggerPrintflnNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	logger := newLogger(LogTrace)
	logger.printfln(LogTrace, "2 + 2 = %d", 4)
}

func TestLoggerPrintflnNoOutputNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	logger := newLogger(LogErrors)
	logger.printfln(LogAll, "2 + 2 = %d", 4)
}

func TestLoggerPrintfNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	logger := newLogger(LogInfo)
	logger.printf(LogInfo, "2 + 2 = %d", 4)
}

func handleLoggerPanic(t *testing.T) {
	handlePanic(t, "Logger")
}
