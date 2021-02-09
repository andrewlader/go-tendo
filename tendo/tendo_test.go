package tendo

import (
	"log"
	"testing"
)

func TestTendoToString(t *testing.T) {
	const testVersion = "0.0.1"
	const targetPath = "../tests/exampletest"

	tendo := NewTendo(LogErrors)
	tendo.version = testVersion
	tendo.Inspect(targetPath, LanguageType(Golang))

	actualTestOutput := tendo.ToString()

	if actualTestOutput != expectedTestOutput {
		t.Errorf("The results of the toString() method did not match the expected results\nExpected:\n%s\n\nActual:\n%s",
			expectedTestOutput, actualTestOutput)
	}
}

func TestTendoDisplayTotals(t *testing.T) {
	const testVersion = "0.0.1"
	const targetPath = "../tests/exampletest"

	defer handlePanic(t, "Tendo")

	tendo := NewTendo(LogErrors)
	tendo.version = testVersion
	tendo.Inspect(targetPath, LanguageType(Golang))

	tendo.DisplayTotals()
}

func TestTendoTestClearSuccess(t *testing.T) {
	const targetPath = "../tests/exampletest"

	tendo := NewTendo(LogErrors)
	tendo.Inspect(targetPath, LanguageType(Golang))

	libraries, _, _, _ := tendo.GetTotals()

	tendo.listener.stop()
	tendo.listener.restart()

	newLibraries, _, _, _ := tendo.GetTotals()

	if (libraries != 1) || (newLibraries != 0) {
		t.Error("Failed to clear out the data in Tendo instance")
	}
}

func TestTendoLibrarySuccess(t *testing.T) {
	const targetPath = "./"
	const expectedLibraries = 4
	const expectedClasses = 6
	const expectedMethods = 26
	const expectedFunctions = 10

	testTendoWithPath(t, LogAll, targetPath, expectedLibraries, expectedClasses, expectedMethods, expectedFunctions)
}

func TestTendoBasicSuccess(t *testing.T) {
	const targetPath = "../tests/exampletest"
	const expectedLibraries = 1
	const expectedClasses = 1
	const expectedMethods = 2
	const expectedFunctions = 1

	testTendoWithPath(t, LogInfo, targetPath, expectedLibraries, expectedClasses, expectedMethods, expectedFunctions)
}

func TestTendoIgnoreTestLibrarySuccess(t *testing.T) {
	const targetPath = "../tests/example_test"
	const expectedLibraries = 0
	const expectedClasses = 0
	const expectedMethods = 0
	const expectedFunctions = 0

	testTendoWithPath(t, LogWarnings, targetPath, expectedLibraries, expectedClasses, expectedMethods, expectedFunctions)
}

func testTendoWithPath(t *testing.T, logLevel LogLevel, targetPath string, expectedLibraries int, expectedClasses int, expectedMethods int, expectedFunctions int) {
	tendo := NewTendo(LogErrors)

	tendo.Inspect(targetPath, LanguageType(Golang))

	libraries, classCount, methodCount, functions := tendo.GetTotals()
	log.Printf("Tendo With Path: %d, %d, %d, %d", libraries, classCount, methodCount, functions)

	if libraries != expectedLibraries {
		t.Errorf("Number of libraries should have been %d, but found %d", expectedLibraries, libraries)
	}
	if classCount != expectedClasses {
		t.Errorf("Number of structs should have been %d, but found %d", expectedClasses, classCount)
	}
	if methodCount != expectedMethods {
		t.Errorf("Number of methods should have been %d, but found %d", expectedMethods, methodCount)
	}
	if functions != expectedFunctions {
		t.Errorf("Number of functions should have been %d, but found %d", expectedFunctions, functions)
	}
}

func handlePanic(t *testing.T, structName string) {
	recovery := recover()
	if recovery != nil {
		t.Errorf("%s function should not panic.", structName)
	}
}

const expectedTestOutput = `

╔╦╗┌─┐┌┐┌┌┬┐┌─┐  ╔╦╗┌─┐┌┬┐┌─┐┬  ┌─┐
 ║ ├┤ │││ │││ │   ║ │ │ │ ├─┤│  └─┐
 ╩ └─┘┘└┘─┴┘└─┘   ╩ └─┘ ┴ ┴ ┴┴─┘└─┘

Version 0.0.1

Source path: ../tests/exampletest

    library exampletest
        class/struct example{}
            method exampleMethod()
            method exampleMethodTwo()
        function exampleFunction()

Totals:
=======
Library Count: 1
Struct Count: 1
Method Count: 2
Function Count: 1
`
