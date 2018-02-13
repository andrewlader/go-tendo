package tendo

import (
	"testing"
)

func TestTendoPackageSuccess(t *testing.T) {
	const targetPath = "./"
	const expectedPackages = 1
	const expectedStructs = 4
	const expectedMethods = 23
	const expectedFunctions = 4

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
