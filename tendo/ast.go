package tendo

import (
	"go/ast"
	"os"
	"path/filepath"
	"strings"
)

// VisitorFunc is used to define the function used when walking trhough a Go source file
type VisitorFunc func(n ast.Node) ast.Visitor

// Visit is used when walking through a Go source file
func (visitor VisitorFunc) Visit(node ast.Node) ast.Visitor {
	return visitor(node)
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
