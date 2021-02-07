package tendo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/andrewlader/go-tendo/tendo/internal/golang"
)

// Tendo is the struct which manages all of the packages in the specified Go project
type Tendo struct {
	languageType   LanguageType
	version        string
	sourcePath     string
	currentPath    string
	currentPackage string
	logger         *Logger
	listener       *listener
	walker         ITendo
}

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
func NewTendo(languageType LanguageType, logLevel LogLevel) *Tendo {
	var walker ITendo

	logger := newLogger(logLevel)
	if logger == nil {
		log.Fatal("Failed to created the logger, so quitting...")
	}

	root := newRoot()
	listener := newListener(root, logger)

	if languageType == LanguageType(Golang) {
		walker = golang.NewGolang(listener.libChan, listener.classChan, listener.methodChan, listener.functionChan)
	}

	return &Tendo{
		version:      "0.0.2",
		languageType: languageType,
		listener:     listener,
		walker:       walker,
		logger:       logger,
	}
}

// Clear clears out all of the data
// func (tendo *Tendo) Clear() {
// 	tendo.packages = make(map[string]*library)
// 	tendo.functions = nil
// }

// Inspect walks through all of the Go files specified in the path and counts the packages, structs and methods
func (tendo *Tendo) Inspect(path string) {
	fullpath, err := filepath.Abs(path)
	if err != nil {
		fullpath = path
	}

	tendo.logger.println(LogAll, asciiArtTendo)
	tendo.logger.printf(LogAll, "### Analysis initiating for path --> %s", fullpath)

	go tendo.listener.Listen()
	tendo.walker.Walk(fullpath)

	folders, err := getListOfFolders(fullpath)
	if err != nil {
		tendo.logger.fatalf(LogErrors, "An error occurred processing the subfolders: %v", err)
	}

	if err == nil {
		for _, path := range folders {
			tendo.walker.Walk(path)
		}
	}

	// all done, so shutdown
	tendo.listener.quitChan <- true
}

func getListOfFolders(path string) ([]string, error) {
	folders := []string{}
	err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isValidFolder(file, path) {
			folders = append(folders, path)
		}
		return nil
	})

	return folders, err
}

func isValidFolder(file os.FileInfo, path string) bool {
	const ignoreHiddenFolders = "."
	const allowCurrentFolder = "./"
	const ignoreVendors = "vendor"
	var ignoreHiddenSubFolders = filepath.ToSlash("/.")

	isValid := false
	if file.IsDir() && !strings.Contains(path, ignoreVendors) && !strings.Contains(path, ignoreHiddenSubFolders) &&
		(path == allowCurrentFolder || !strings.HasPrefix(path, ignoreHiddenFolders)) {
		isValid = true
	}

	return isValid
}

// DisplayTotals calls GetTotals() and then displays the results to the console
func (tendo *Tendo) DisplayTotals() {
	tendo.logger.println(logAlways, tendo.ToString())
}

// GetTotals returns the total number of packages, structs and methods
func (tendo *Tendo) GetTotals() (int, int, int, int) {
	structCount := 0
	methodCount := 0
	functionCount := 0

	for _, lib := range tendo.listener.root.libraries {
		structCount += len(lib.classes)
		functionCount += len(lib.functions)
		for _, class := range lib.classes {
			methodCount += len(class.methods)
		}
	}

	return len(tendo.listener.root.libraries), structCount, methodCount, functionCount
}

func (tendo *Tendo) ToString() string {
	const indent = "    "

	outputPrefix := fmt.Sprintf("%s%s\n\nSource path: %s\n", asciiArtTendoTotals, tendo.version, tendo.sourcePath)

	var tree []string

	// for each of the packages
	for _, pkg := range tendo.listener.root.libraries {
		tree = append(tree, fmt.Sprintf("%spackage %s", indent, pkg.name))
		// display all the structs in the package
		for _, class := range pkg.classes {
			tree = append(tree, fmt.Sprintf("%s%sclass/struct %s{}", indent, indent, class.name))
			// and display all of the methods for the structs
			for _, method := range class.methods {
				tree = append(tree, fmt.Sprintf("%s%s%smethod %s()", indent, indent, indent, method))
			}
		}

		// and display the functions in the package
		for _, function := range pkg.functions {
			tree = append(tree, fmt.Sprintf("%s%sfunction %s()", indent, indent, function))
		}
	}

	packages, structCount, methodCount, functions := tendo.GetTotals()

	tree = append(tree, fmt.Sprintf("\nTotals:\n=======\nPackage Count: %d\nStruct Count: %d\nMethod Count: %d\nFunction Count: %d\n",
		packages, structCount, methodCount, functions))

	output := fmt.Sprintf("%s\n%s", outputPrefix, strings.Join(tree[:], "\n"))

	return output
}
