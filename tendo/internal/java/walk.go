package java

import (
	"go/ast"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/andrewlader/go-tendo/tendo/internal"
)

func Walk(path string) {}

type Java struct {
	pkgChan      chan *internal.LibraryInfo
	classChan    chan *internal.ClassInfo
	methodChan   chan *internal.MethodInfo
	functionChan chan *internal.FunctionInfo
}

func NewJava(
	pkgChan chan *internal.LibraryInfo,
	classChan chan *internal.ClassInfo,
	methodChan chan *internal.MethodInfo,
	functionChan chan *internal.FunctionInfo) *Java {
	return &Java{
		pkgChan:      pkgChan,
		classChan:    classChan,
		methodChan:   methodChan,
		functionChan: functionChan,
	}
}

func (java *Java) Walk(path string) {
	const suffixJava = ".java"

	log.Printf("## Inspecting path --> %s", path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Skipping path, failed to parse path: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), suffixJava) {
				pkg := &ast.Package{
					Name: path,
				}
				java.inspectPackage(pkg)
				java.inspectFile(path, file)
			}
		}
	}
}

func (java *Java) inspectFile(pathToFile string, file os.FileInfo) {
	log.Printf("### Inspecting file --> %s", file.Name())

	inputs := []string{pathToFile, file.Name()}
	filename := strings.Join(inputs, "\\")
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("error reading file \"%s\": %v", file.Name(), err)
		return
	}

	fileSource := string(bytes)
	lines := strings.Split(fileSource, "\n")

	pkgName := java.handlePackage(pathToFile)
	java.examineFileLineByLine(file.Name(), pkgName, lines)
}

func (java *Java) examineFileLineByLine(filename string, pkgName string, lines []string) {
	const matchClassNames = "(class)[\\s]+([a-zA-Z0-9]+)[\\s]*{"
	const matchFunctionNames = "((public)|(private))\\s((static)\\s)?([a-zA-Z0-9]+[\\s]+)(([a-zA-Z0-9]+)\\(.*\\)[\\s]*{)"
	const matchBraces = "({)|(})[\\s]*"

	regexClassNames := regexp.MustCompile(matchClassNames)
	regexFunctionNames := regexp.MustCompile(matchFunctionNames)
	regexBraces := regexp.MustCompile(matchBraces)
	braceCount := 0
	stack := &stackWalkState{}
	currentClassName := ""

	for _, line := range lines {
		classNames := regexClassNames.FindStringSubmatch(line)
		lenClassNames := len(classNames) - 1
		functionNames := regexFunctionNames.FindStringSubmatch(line)
		lenFunctionNames := len(functionNames) - 1
		braces := regexBraces.FindStringSubmatch(line)
		lenBraces := len(braces) - 1

		if lenClassNames > 0 {
			currentClassName = classNames[lenClassNames]
			java.handleClass(filename, pkgName, currentClassName)
			stack.push(withinClass)
			// found a class, so start counting braces at 1 since there was one in that line
			braceCount = 1
		} else if lenFunctionNames > 0 {
			currentFunctionName := functionNames[lenFunctionNames]

			if stack.peek() == withinClass {
				java.handleMethod(currentClassName, currentFunctionName)
			} else {
				java.handleFunction(currentFunctionName)
			}
			stack.push(withinMethodOrFunction)
			// found a function/method, so start counting braces at 1 since there was one in that line
			braceCount = 1
		} else if lenBraces > 0 {
			if braces[0] == "{" {
				braceCount++
			} else {
				braceCount--
				if braceCount == 0 {
					stack.pop()
					if stack.peek() != withinNothing {
						// not within nothing, so reset braces to 1
						braceCount = 1
					}
				}
			}
		}
	}
}

func (java *Java) handlePackage(pathToFile string) string {
	paths := strings.Split(pathToFile, "\\")
	numberOfPaths := len(paths)
	parentPath := ""
	thisPath := ""
	if numberOfPaths > 0 {
		if numberOfPaths > 1 {
			parentPath = paths[numberOfPaths-2]
		}
		thisPath = paths[numberOfPaths-1]
		pkgInfo := &internal.LibraryInfo{
			Parent: parentPath,
			Name:   thisPath,
		}
		java.pkgChan <- pkgInfo
	}

	return thisPath
}

func (java *Java) handleClass(filename string, pkgName string, className string) {
	classInfo := &internal.ClassInfo{
		Package: pkgName,
		File:    filename,
		Name:    className,
	}
	java.classChan <- classInfo
}

func (java *Java) handleMethod(className string, methodName string) {
	methodInfo := &internal.MethodInfo{
		Class: className,
		Name:  methodName,
	}
	java.methodChan <- methodInfo
}

func (java *Java) handleFunction(functionName string) {
	functionInfo := &internal.FunctionInfo{
		Name: functionName,
	}
	java.functionChan <- functionInfo
}

func (java *Java) inspectPackage(pkg *ast.Package) bool {

	return true
}

func (java *Java) inspectFunction(function *ast.FuncDecl) {
}
