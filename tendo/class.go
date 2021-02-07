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

func (obj *Class) addMethod(name string, logger *Logger) {
	methodFound := false
	for _, existingMethod := range obj.methods {
		if existingMethod == name {
			methodFound = true
		}
	}

	if !methodFound {
		obj.methods = append(obj.methods, name)
		logger.printf(LogTrace, "Added method %s to class %s", name, obj.name)
	}
}

func (obj *Class) GetMethodCount() int {
	return len(obj.methods)
}
