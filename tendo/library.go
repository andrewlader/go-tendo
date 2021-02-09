package tendo

type library struct {
	name      string
	classes   map[string]*Class
	functions []string
	libraries map[string]*library
}

func newLibrary(name string) *library {
	return &library{
		name:      name,
		classes:   make(map[string]*Class),
		libraries: make(map[string]*library),
	}
}

func (lib *library) addClass(name string, theLogger *Logger) {
	_, ok := lib.classes[name]
	if !ok {
		obj := newClass(name)
		lib.classes[name] = obj
		theLogger.printf(LogTrace, "Added class %s to package %s", name, lib.name)
	}
}

func (lib *library) addLibrary(name string, theLogger *Logger) {
	_, ok := lib.libraries[name]
	if !ok {
		lib.libraries[name] = newLibrary(name)
		theLogger.printf(LogTrace, "Added package %s to package %s", name, lib.name)
	}
}

func (lib *library) addFunction(name string, theLogger *Logger) {
	funcFound := false
	for _, existingFunc := range lib.functions {
		if existingFunc == name {
			funcFound = true
		}
	}

	if !funcFound {
		lib.functions = append(lib.functions, name)
		theLogger.printf(LogTrace, "Added function %s to package %s", name, lib.name)
	}
}

func (lib *library) getSubPackageCount() int {
	return len(lib.libraries)
}

func (lib *library) getClassCount() int {
	return len(lib.classes)
}
