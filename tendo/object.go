package tendo

import "fmt"

type object struct {
	name    string
	methods []string
}

func newObject(name string) *object {
	return &object{
		name: name,
	}
}

func (obj *object) addMethod(name string) error {
	methodFound := false
	for _, existingMethod := range obj.methods {
		if existingMethod == name {
			methodFound = true
		}
	}

	if !methodFound {
		obj.methods = append(obj.methods, name)
		return nil
	}

	return fmt.Errorf("Trace Warning: attempted to add a method named '%s' to struct '%s' when it already existed", name, obj.name)
}

func (obj *object) getMethodCount() int {
	return len(obj.methods)
}
