package tendo

import "testing"

const (
	expectedPackageName         = "simple_package"
	expectedPackageObjectName   = "simplePackageObject"
	expectedPackageFunctionName = "simplePackageFunction"
)

func TestPackageAddObjectSuccess(t *testing.T) {
	pkg := newLibrary(expectedPackageName)

	pkg.addClass(expectedPackageObjectName, theLogger)

	if pkg.name != expectedPackageName {
		t.Errorf("Failed to return proper pkg.name. Expected '%s', but got '%s'", expectedPackageName, pkg.name)
	}
	if pkg.getClassCount() != 1 {
		t.Errorf("Failed to return proper number of objects. Expected 1, but got %d", pkg.getClassCount())
	}
	if (pkg.getClassCount() == 1) && (pkg.classes[expectedPackageObjectName].name != expectedPackageObjectName) {
		t.Errorf("Failed to return proper object name. Expected '%s', but got '%s'", expectedPackageObjectName, pkg.classes[expectedPackageObjectName].name)
	}
}

func TestPackageAddSameObjectError(t *testing.T) {
	defer handlePanic(t, "library")

	pkg := newLibrary(expectedPackageName)

	pkg.addClass(expectedPackageObjectName, theLogger)

	// duplicate call should cause no panic or other issues
	pkg.addClass(expectedPackageObjectName, theLogger)
}

func TestPackageAddFunctionSuccess(t *testing.T) {
	pkg := newLibrary(expectedPackageName)

	pkg.addFunction(expectedPackageFunctionName, theLogger)

	if len(pkg.functions) != 1 {
		t.Errorf("Failed to return proper number of functions. Expected 1, but got %d", len(pkg.functions))
	}
	if (len(pkg.functions) == 1) && (pkg.functions[0] != expectedPackageFunctionName) {
		t.Errorf("Failed to return proper function name. Expected '%s', but got '%s'", expectedPackageFunctionName, pkg.functions[0])
	}
}

func TestPackageAddSameFunctionError(t *testing.T) {
	defer handlePanic(t, "library")

	pkg := newLibrary(expectedPackageName)

	pkg.addFunction(expectedPackageFunctionName, theLogger)

	// duplicate call should cause no panic or other issues
	pkg.addFunction(expectedPackageFunctionName, theLogger)
}
