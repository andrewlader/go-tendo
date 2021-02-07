package tendo

import "testing"

func TestObjectAddMethodSuccess(t *testing.T) {
	const expectedObjectName = "simpleObject"
	const expectedObjectMethodName = "simpleObjectMethod"

	obj := Class{
		name: expectedObjectName,
	}

	obj.addMethod(expectedObjectMethodName)

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

func TestObjectAddSameMethodError(t *testing.T) {
	const expectedObjectName = "simpleObject"
	const expectedObjectMethodName = "simpleObjectMethod"

	obj := Class{
		name: expectedObjectName,
	}

	obj.addMethod(expectedObjectMethodName)

	err := obj.addMethod(expectedObjectMethodName)
	if err == nil {
		t.Error("Failed to get an error when adding the same method name to the same object")
	}
}
