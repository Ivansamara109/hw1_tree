package main

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"io/ioutil"
)

var disabledFound = []string{".git", ".gitignore", ".idea", "README.md", "."}

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
	var pathResult string
	var tabs string

	path = strings.Replace(path, `\`, `/`, 100)
	pathListFull := strings.Split(path, `/`)
	pathList := pathListFull[1:]

	if len(pathList) == 0 {
		return pathResult
	}

	basePath := filepath.Base(path)
	if isLastElementPath(path) {
		pathResult = pathResult + `└───` + basePath
	} else {
		pathResult = pathResult + `├───` + basePath
	}


	tabs = getTabFormat(pathListFull)

	return tabs + pathResult + "\n"
}

func getTabFormat(pathList []string) string {
	var tabResult string

	for i := 2; i < len(pathList); i++ {
		if isLastElementPath(filepath.Join(pathList[:i]...)) {
			tabResult = tabResult + "\t"
		} else {
			tabResult = tabResult + "│\t"
		}
	}

	return tabResult
}

func isLastElementPath(path string) bool  {

	basePath := filepath.Base(path)

	var catalogList []string
	var fileList []string

	files, _ := ioutil.ReadDir(filepath.Dir(path))

	for _, file := range files {
		if file.IsDir() {
			catalogList = append(catalogList, file.Name())
		} else {
			fileList = append(fileList, file.Name())
		}
	}

	if basePath == catalogList[len(catalogList)-1] || basePath == fileList[len(fileList)-1]{
		return true
	}

	return false
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