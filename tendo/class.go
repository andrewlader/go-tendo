package tendo

type Class struct {
	name    string
	methods []string
}

func newClass(name string) *Class {
	return &Class{
		name: name,
	}
}

func (obj *Class) addMethod(name string, theLogger *Logger) {
	methodFound := false
	for _, existingMethod := range obj.methods {
		if existingMethod == name {
			methodFound = true
		}
	}

	if !methodFound {
		obj.methods = append(obj.methods, name)
		theLogger.printf(LogTrace, "Added method %s to class %s", name, obj.name)
	}
}

func (obj *Class) getMethodCount() int {
	return len(obj.methods)
}
