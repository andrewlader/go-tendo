package tendo

import "testing"

const (
	expectedPackageName         = "simple_package"
	expectedPackageObjectName   = "simplePackageObject"
	expectedPackageFunctionName = "simplePackageFunction"
)

func TestPackageAddObjectSuccess(t *testing.T) {
	pkg := newPackage(expectedPackageName)

	pkg.addObject(expectedPackageObjectName)

	if pkg.name != expectedPackageName {
		t.Errorf("Failed to return proper pkg.name. Expected '%s', but got '%s'", expectedPackageName, pkg.name)
	}
	if pkg.getObjectCount() != 1 {
		t.Errorf("Failed to return proper number of objects. Expected 1, but got %d", pkg.getObjectCount())
	}
	if (pkg.getObjectCount() == 1) && (pkg.objects[expectedPackageObjectName].name != expectedPackageObjectName) {
		t.Errorf("Failed to return proper object name. Expected '%s', but got '%s'", expectedPackageObjectName, pkg.objects[expectedPackageObjectName].name)
	}
}

func TestPackageAddSameObjectError(t *testing.T) {
	pkg := newPackage(expectedPackageName)

	pkg.addObject(expectedPackageObjectName)

	err := pkg.addObject(expectedPackageObjectName)
	if err == nil {
		t.Error("Failed to get an error when adding the same object name to the same package")
	}
}

func TestPackageAddFunctionSuccess(t *testing.T) {
	pkg := newPackage(expectedPackageName)

	pkg.addFunction(expectedPackageFunctionName)

	if len(pkg.functions) != 1 {
		t.Errorf("Failed to return proper number of functions. Expected 1, but got %d", len(pkg.functions))
	}
	if (len(pkg.functions) == 1) && (pkg.functions[0] != expectedPackageFunctionName) {
		t.Errorf("Failed to return proper function name. Expected '%s', but got '%s'", expectedPackageFunctionName, pkg.functions[0])
	}
}

func TestPackageAddSameFunctionError(t *testing.T) {
	pkg := newPackage(expectedPackageName)

	pkg.addFunction(expectedPackageFunctionName)

	err := pkg.addFunction(expectedPackageFunctionName)
	if err == nil {
		t.Error("Failed to get an error when adding the same function name to the same package")
	}
}
