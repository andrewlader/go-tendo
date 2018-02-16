package tendo

import (
	"testing"
)

func TestTendoToString(t *testing.T) {
	const testVersion = "0.0.1"
	const targetPath = "../tests/exampletest"

	tendo := NewTendo(LogErrors)
	tendo.version = testVersion
	tendo.Inspect(targetPath)

	actualTestOutput := tendo.toString()

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
	tendo.Inspect(targetPath)

	tendo.DisplayTotals()
}

func TestTendoTestClearSuccess(t *testing.T) {
	const targetPath = "../tests/exampletest"

	tendo := NewTendo(LogErrors)
	tendo.Inspect(targetPath)

	packages, _, _, _ := tendo.GetTotals()

	tendo.Clear()

	newPackages, _, _, _ := tendo.GetTotals()

	if (packages != 1) || (newPackages != 0) {
		t.Error("Failed to clear out the data in Tendo instance")
	}
}

func TestTendoPackageSuccess(t *testing.T) {
	const targetPath = "./"
	const expectedPackages = 1
	const expectedStructs = 4
	const expectedMethods = 23
	const expectedFunctions = 6

	testTendoWithPath(t, LogAll, targetPath, expectedPackages, expectedStructs, expectedMethods, expectedFunctions)
}

func TestTendoBasicSuccess(t *testing.T) {
	const targetPath = "../tests/exampletest"
	const expectedPackages = 1
	const expectedStructs = 1
	const expectedMethods = 2
	const expectedFunctions = 1

	testTendoWithPath(t, LogInfo, targetPath, expectedPackages, expectedStructs, expectedMethods, expectedFunctions)
}

func TestTendoIgnoreTestPackageSuccess(t *testing.T) {
	const targetPath = "../tests/example_test"
	const expectedPackages = 0
	const expectedStructs = 0
	const expectedMethods = 0
	const expectedFunctions = 0

	testTendoWithPath(t, LogWarnings, targetPath, expectedPackages, expectedStructs, expectedMethods, expectedFunctions)
}

func testTendoWithPath(t *testing.T, logLevel LogLevel, targetPath string, expectedPackages int, expectedStructs int, expectedMethods int, expectedFunctions int) {
	tendo := NewTendo(LogErrors)

	tendo.Inspect(targetPath)

	packages, structCount, methodCount, functions := tendo.GetTotals()

	if packages != expectedPackages {
		t.Errorf("Number of packages should have been %d, but found %d", expectedPackages, packages)
	}
	if structCount != expectedStructs {
		t.Errorf("Number of structs should have been %d, but found %d", expectedStructs, structCount)
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

    package exampletest
        struct example{}
            method exampleMethod()
            method exampleMethodTwo()
        function exampleFunction()

Totals:
=======
Package Count: 1
Struct Count: 1
Method Count: 2
Function Count: 1
`
