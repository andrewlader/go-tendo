package tendo

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Tendo is the struct which manages all of the packages in the specified Go project
type Tendo struct {
	sourcePath     string
	currentPath    string
	logger         *logger
	currentPackage string
	packages       map[string]*pkg
	functions      []string
}

const version = "0.0.1"

const asciiArtTendoTotals = `

╔╦╗┌─┐┌┐┌┌┬┐┌─┐  ╔╦╗┌─┐┌┬┐┌─┐┬  ┌─┐
 ║ ├┤ │││ │││ │   ║ │ │ │ ├─┤│  └─┐
 ╩ └─┘┘└┘─┴┘└─┘   ╩ └─┘ ┴ ┴ ┴┴─┘└─┘

Version `

const asciiArtTendo = `

╔╦╗┌─┐┌┐┌┌┬┐┌─┐
 ║ ├┤ │││ │││ │
 ╩ └─┘┘└┘─┴┘└─┘

 `

// NewTendo creates a new instance of Tendo and returns a reference to it
func NewTendo(logLevel LogLevel) *Tendo {
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	logger := newLogger(logLevel)

	return &Tendo{
		currentPath: currentPath,
		logger:      logger,
		packages:    make(map[string]*pkg),
	}
}

// Clear clears out all of the data
func (tendo *Tendo) Clear() {
	tendo.packages = make(map[string]*pkg)
	tendo.functions = nil
}

// DisplayTotals calls GetTotals() and then displays the results to the console
func (tendo *Tendo) DisplayTotals() {
	tendo.logger.printfln(logAlways, "%s%s", asciiArtTendoTotals, version)
	tendo.logger.printfln(logAlways, "\nSource path: %s", tendo.sourcePath)

	tendo.displayTree()

	packages, structCount, methodCount, functions := tendo.GetTotals()

	tendo.logger.printfln(logAlways, "\nTotals:\n=======\nPackage Count: %d\nStruct Count: %d\nMethod Count: %d\nFunction Count: %d\n",
		packages, structCount, methodCount, functions)

}

// GetTotals returns the total number of packages, structs and methods
func (tendo *Tendo) GetTotals() (int, int, int, int) {
	structCount := 0
	methodCount := 0
	functionCount := 0

	for _, pkg := range tendo.packages {
		structCount += pkg.getObjectCount()
		functionCount += len(pkg.functions)
		for _, obj := range pkg.objects {
			methodCount += obj.getMethodCount()
		}
	}

	return len(tendo.packages), structCount, methodCount, functionCount
}

// Inspect walks through all of the Go files specified in the path and counts the packages, structs and methods
func (tendo *Tendo) Inspect(path string) {
	tendo.sourcePath = path

	fullpath, err := filepath.Abs(path)
	if err != nil {
		fullpath = path
	}

	tendo.logger.println(LogAll, asciiArtTendo)
	tendo.logger.printf(LogAll, "### Analysis initiating for path --> %s", path)

	folders, err := getListOfFolders(fullpath)
	if err != nil {
		tendo.logger.fatalf(LogErrors, "An error occurred processing the subfolders: %v", err)
	}

	if err == nil {
		for _, path := range folders {
			tendo.inspectFolder(path)
		}
	}
}

func (tendo *Tendo) inspectFolder(path string) {
	fileSet := token.NewFileSet()

	relativePath, err := filepath.Rel(tendo.currentPath, path)
	if err != nil {
		tendo.logger.printf(LogTrace, "trace warning: %v", err)
		relativePath = path
	}
	tendo.logger.printf(LogTrace, "## Inspecting path --> %s", relativePath)

	pkgs, err := parser.ParseDir(fileSet, path, nil, 0)
	if err != nil {
		tendo.logger.printf(LogWarnings, "Skipping path, failed to parse path: %v", err)
	}

	for _, pkg := range pkgs {
		ast.Walk(VisitorFunc(tendo.inspectNode), pkg)
	}
}

func (tendo *Tendo) inspectNode(node ast.Node) ast.Visitor {
	switch nodeType := node.(type) {
	case *ast.Package:
		if tendo.inspectPackage(nodeType) {
			return VisitorFunc(tendo.inspectNode)
		}

	case *ast.File:
		return VisitorFunc(tendo.inspectNode)

	case *ast.FuncDecl:
		tendo.inspectFunction(nodeType)
	}

	return nil
}

func (tendo *Tendo) inspectFunction(function *ast.FuncDecl) {
	if function.Recv != nil {
		field := function.Recv.List[0]
		if receiver, ok := field.Type.(*ast.StarExpr); ok {
			structName := fmt.Sprintf("%s", receiver.X)
			tendo.addStruct(structName)
			tendo.addMethod(structName, function.Name.Name)
		}
	} else {
		tendo.logger.printf(LogTrace, "Added function --> %s", function.Name.Name)
		tendo.addFunction(function.Name.Name)
	}
}

func (tendo *Tendo) inspectPackage(pkg *ast.Package) bool {
	const ignoreTestPackages = "_test"

	packageName := pkg.Name
	if strings.HasSuffix(packageName, ignoreTestPackages) {
		tendo.logger.printf(LogTrace, "skipping package --> %s", packageName)
		return false
	}

	pkg = tendo.pruneTestFiles(pkg)
	tendo.currentPackage = packageName
	tendo.addPackage(packageName)
	return true
}

func (tendo *Tendo) pruneTestFiles(pkg *ast.Package) *ast.Package {
	const ignoreTestFiles = "_test.go"

	// prune off the test packages
	listOfTestFiles := []string{}
	for filename := range pkg.Files {
		if strings.HasSuffix(filename, ignoreTestFiles) {
			listOfTestFiles = append(listOfTestFiles, filename)
		}
	}
	for _, filename := range listOfTestFiles {
		tendo.logger.printf(LogTrace, "Skipping test file --> %s", filename)
		delete(pkg.Files, filename)
	}

	return pkg
}

func (tendo *Tendo) addPackage(name string) {
	_, ok := tendo.packages[name]
	if !ok {
		newPkg := &pkg{
			name:    name,
			objects: make(map[string]*object),
		}
		tendo.packages[name] = newPkg
		tendo.logger.printf(LogTrace, "Added package --> %s", name)
	}
}

func (tendo *Tendo) addStruct(structName string) {
	err := tendo.packages[tendo.currentPackage].addObject(structName)
	if err != nil {
		tendo.logger.printf(LogTrace, "%s", err)
	} else {
		tendo.logger.printf(LogTrace, "Added struct --> %s", structName)
	}
}

func (tendo *Tendo) addFunction(name string) {
	err := tendo.packages[tendo.currentPackage].addFunction(name)
	if err != nil {
		tendo.logger.printf(LogTrace, "%s", err)
	} else {
		tendo.logger.printf(LogTrace, "Added function --> %s", name)
	}
}

func (tendo *Tendo) addMethod(structName string, methodName string) {
	err := tendo.packages[tendo.currentPackage].objects[structName].addMethod(methodName)

	if err != nil {
		tendo.logger.printf(LogTrace, "%s", err)
	} else {
		tendo.logger.printf(LogTrace, "Added method --> %s", methodName)
	}
}

func (tendo *Tendo) displayTree() {
	const indent = "    "

	tendo.logger.println(LogInfo, "\n")

	// for each of the packages
	for _, pkg := range tendo.packages {
		tendo.logger.printfln(LogInfo, "%spackage %s", indent, pkg.name)
		// display all the structs in the package
		for _, object := range pkg.objects {
			tendo.logger.printfln(LogInfo, "%s%sstruct %s", indent, indent, object.name)
			// and display all of the methods for the structs
			for _, method := range object.methods {
				tendo.logger.printfln(LogInfo, "%s%s%smethod %s()", indent, indent, indent, method)
			}
		}

		// and display the functions in the package
		for _, function := range pkg.functions {
			tendo.logger.printfln(LogInfo, "%s%sfunction %s()", indent, indent, function)
		}
	}
}
