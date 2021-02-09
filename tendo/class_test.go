package tendo

import (
	"testing"
)

var theLogger = &Logger{
	level: LogLevel(LogAll),
}

func TestClassAddMethodSuccess(t *testing.T) {
	const expectedObjectName = "simpleObject"
	const expectedObjectMethodName = "simpleObjectMethod"

	obj := Class{
		name: expectedObjectName,
	}

	obj.addMethod(expectedObjectMethodName, theLogger)

	if obj.name != expectedObjectName {
		t.Errorf("Failed to return proper object.name. Expected '%s', but got '%s'", expectedObjectName, obj.name)
	}
	if obj.getMethodCount() != 1 {
		t.Errorf("Failed to return proper number of methods. Expected 1, but got %d", obj.getMethodCount())
	}
	if (obj.getMethodCount() == 1) && (obj.methods[0] != expectedObjectMethodName) {
		t.Errorf("Failed to return proper method name. Expected '%s', but got '%s'", expectedObjectMethodName, obj.methods[0])
	}
}

func TestClassAddSameMethodError(t *testing.T) {
	const expectedObjectName = "simpleObject"
	const expectedObjectMethodName = "simpleObjectMethod"

	defer handlePanic(t, "class")

	obj := Class{
		name: expectedObjectName,
	}

	obj.addMethod(expectedObjectMethodName, theLogger)

	// duplicate call should cause no panic or other issues
	obj.addMethod(expectedObjectMethodName, theLogger)
}
