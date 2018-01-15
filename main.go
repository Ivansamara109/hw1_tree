package main

import (
	//"fmt"
	//"io"
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

var disabledFound = []string{".git", ".gitignore", ".idea", "README.md", "."}


const (
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out interface{}, filePath string, printFiles bool) error  {

	var resultTree string

	fileList, err := getFileList(filePath, printFiles)

	for _, file := range fileList {
		resultTree = resultTree + formatPath(file)
	}


	fmt.Println(resultTree)

	return err
}


func formatPath(path string) string {
	path = strings.Replace(path, `\`, `/`, 100)
	var pathResult string

	pathListFull := strings.Split(path, `/`)
	pathList := pathListFull[1:]

	if len(pathList) == 0 {
		return pathResult
	}

	for index, item := range pathList {
		if index == (len(pathList) - 1) {
			pathResult = pathResult + `├───` + item
		} else {
			pathResult = pathResult +  "│\t"
		}
	}

	return pathResult + "\n"
}


func getFileList(filePath string, printFiles bool) ([]string, error)  {
	var fileList []string

	err := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		if isDisabled(path) {
			return nil
		}

		if !f.IsDir() && !printFiles {
			return nil
		}

		fileList = append(fileList, path)
		return nil
	})

	return fileList, err
}


func isDisabled(path string) bool{
	pathList := strings.Split(path, `\`)

	for _, value := range disabledFound {
		if pathList[0] == value {
			return true
		}
	}
	return false
}