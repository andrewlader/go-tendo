package internal

type LibraryInfo struct {
	Parent string
	Name   string
}

type ClassInfo struct {
	Package string
	File    string
	Name    string
}

type MethodInfo struct {
	Class string
	Name  string
}

type FunctionInfo struct {
	File string
	Name string
}
