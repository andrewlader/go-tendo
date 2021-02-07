package golang

import (
	"go/ast"
)

// VisitorFunc is used to define the function used when walking through a Go source file
type VisitorFunc func(node ast.Node) ast.Visitor

// Visit is used when walking through a Go source file
func (visitor VisitorFunc) Visit(node ast.Node) ast.Visitor {
	return visitor(node)
}
