package tendo

import (
	"fmt"
)

type pkg struct {
	name      string
	objects   map[string]*object
	functions []string
}

func newPackage(name string) *pkg {
	return &pkg{
		name:    name,
		objects: make(map[string]*object),
	}
}

func (pkg *pkg) addObject(name string) error {
	_, ok := pkg.objects[name]
	if !ok {
		obj := newObject(name)
		pkg.objects[name] = obj
		return nil
	}

	return fmt.Errorf("Trace Warning: attempted to add a struct named '%s' to package '%s' when it already existed", name, pkg.name)
}

func (pkg *pkg) addFunction(name string) error {
	funcFound := false
	for _, existingFunc := range pkg.functions {
		if existingFunc == name {
			funcFound = true
		}
	}

	if !funcFound {
		pkg.functions = append(pkg.functions, name)
		return nil
	}

	return fmt.Errorf("Trace Warning: attempted to add a function named '%s' to package '%s' when it already existed", name, pkg.name)
}
func (pkg *pkg) getObjectCount() int {
	return len(pkg.objects)
}
