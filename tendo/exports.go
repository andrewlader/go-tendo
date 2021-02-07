package tendo

// LanguageType determines which language the repo uses
type LanguageType uint8

// The various types of languages
const (
	Golang LogLevel = iota
	Java
)

// NodeType determines which node was encountered when walking the path
type NodeType uint8

// The various types of nodes
const (
	NoType NodeType = iota
	PackageType
	ClassType
	MethodType
	FunctionType
)

type ITendo interface {
	Walk(path string)
}
