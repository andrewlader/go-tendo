package tendo

import "testing"

func TestLoggerPrintlnNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	theLogger := newLogger(LogInfo)
	theLogger.printfln(LogInfo, "2 + 2 = 4")
}

func TestLoggerPrintflnNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	theLogger := newLogger(LogTrace)
	theLogger.printfln(LogTrace, "2 + 2 = %d", 4)
}

func TestLoggerPrintflnNoOutputNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	theLogger := newLogger(LogErrors)
	theLogger.printfln(LogAll, "2 + 2 = %d", 4)
}

func TestLoggerPrintfNoPanic(t *testing.T) {
	defer handleLoggerPanic(t)

	theLogger := newLogger(LogInfo)
	theLogger.printf(LogInfo, "2 + 2 = %d", 4)
}

func handleLoggerPanic(t *testing.T) {
	handlePanic(t, "theLogger")
}
