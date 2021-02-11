package golang

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/andrewlader/go-tendo/tendo/internal"
)

type Golang struct {
	pkgChan      chan *internal.LibraryInfo
	classChan    chan *internal.ClassInfo
	methodChan   chan *internal.MethodInfo
	functionChan chan *internal.FunctionInfo
}

func NewGolang(
	pkgChan chan *internal.LibraryInfo,
	classChan chan *internal.ClassInfo,
	methodChan chan *internal.MethodInfo,
	functionChan chan *internal.FunctionInfo) *Golang {
	return &Golang{
		pkgChan:      pkgChan,
		classChan:    classChan,
		methodChan:   methodChan,
		functionChan: functionChan,
	}
}

func (golang *Golang) Walk(path string) {
	fileSet := token.NewFileSet()

	log.Printf("## Inspecting path --> %s", path)

	pkgs, err := parser.ParseDir(fileSet, path, nil, 0)
	if err != nil {
		log.Printf("Skipping path, failed to parse path: %v", err)
	}

	for _, pkg := range pkgs {
		ast.Walk(VisitorFunc(golang.inspectNode), pkg)
	}
}

func (golang *Golang) inspectNode(node ast.Node) ast.Visitor {
	switch nodeType := node.(type) {
	case *ast.Package:
		if golang.inspectPackage(nodeType) {
			return VisitorFunc(golang.inspectNode)
		}

	case *ast.File:
		// ignore files for now...
		return VisitorFunc(golang.inspectNode)

	case *ast.FuncDecl:
		golang.inspectFunction(nodeType)
	}

	return nil
}

func (golang *Golang) inspectPackage(pkg *ast.Package) bool {
	const ignoreTestPackages = "_test"

	packageName := pkg.Name
	if strings.HasSuffix(packageName, ignoreTestPackages) {
		return false
	}

	pkg = golang.pruneTestFiles(pkg)

	pkgInfo := &internal.LibraryInfo{
		Name: packageName,
	}
	golang.pkgChan <- pkgInfo

	return true
}

func (golang *Golang) inspectFunction(function *ast.FuncDecl) {
	if function.Recv != nil {
		field := function.Recv.List[0]
		if receiver, ok := field.Type.(*ast.StarExpr); ok {
			className := fmt.Sprintf("%s", receiver.X)
			methodInfo := &internal.MethodInfo{
				Class: className,
				Name:  function.Name.Name,
			}
			golang.methodChan <- methodInfo
		}
	} else {
		log.Printf("Added function --> %s", function.Name.Name)
		functionInfo := &internal.FunctionInfo{
			Name: function.Name.Name,
		}
		golang.functionChan <- functionInfo
	}
}

func (golang *Golang) pruneTestFiles(pkg *ast.Package) *ast.Package {
	const ignoreTestFiles = "_test.go"

	// prune off the test packages
	listOfTestFiles := []string{}
	for filename := range pkg.Files {
		if strings.HasSuffix(filename, ignoreTestFiles) {
			listOfTestFiles = append(listOfTestFiles, filename)
		}
	}
	for _, filename := range listOfTestFiles {
		log.Printf("Skipping test file --> %s", filename)
		delete(pkg.Files, filename)
	}

	return pkg
}
