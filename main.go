package main

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"io/ioutil"
	"io"
)

var disabledFound = []string{".git", ".gitignore", ".idea", "README.md", ".", "test_compare"}
var printFilesGlobal bool


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

func dirTree(out io.Writer, filePath string, printFiles bool) error  {

	var resultTree string
	printFilesGlobal = printFiles

	fileList, err := getFileList(filePath)

	for _, file := range fileList {
		resultTree = resultTree + formatPath(file)
	}

	fmt.Fprintln(out, resultTree)

	return err
}

func getFileList(filePath string) ([]string, error)  {
	var fileList []string

	err := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		if isDisabled(path) {
			return nil
		}

		if !f.IsDir() && !printFilesGlobal {
			return nil
		}

		fileList = append(fileList, path)
		return nil
	})

	return fileList, err
}


func formatPath(pathOrigin string) string {
	var pathResult string
	var tabs string

	pathLinux := strings.Replace(pathOrigin, `\`, `/`, 100)
	pathListFull := strings.Split(pathLinux, `/`)
	pathList := pathListFull[1:]

	if len(pathList) == 0 {
		return pathResult
	}

	basePath := fmt.Sprintf("%s %s", filepath.Base(pathOrigin), getFileSize(pathOrigin))

	if isLastElementPath(pathOrigin) {
		pathResult = pathResult + `└───` + basePath
	} else {
		pathResult = pathResult + `├───` + basePath
	}

	tabs = getTabs(pathListFull)

	return tabs + pathResult + "\n"
}

//Формат отступов для дерева
func getTabs(pathList []string) string {
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

//Проверка на то что элемент последний в списке
func isLastElementPath(path string) bool  {

	basePath := filepath.Base(path)

	var sortList []string

	files, _ := ioutil.ReadDir(filepath.Dir(path))

	for _, file := range files {
		if printFilesGlobal == false && file.IsDir() == false {
			continue
		}
		sortList = append(sortList, file.Name())
	}

	if sortList[len(sortList)-1] == basePath {
		return true
	}

	return false
}

//Размер файла
func getFileSize(path string) string  {
	var fileSize string
	fileInfo, _ := os.Stat(path)
	if !fileInfo.IsDir() {
		size := fileInfo.Size()
		if size == 0 {
			fileSize = "(empty)"
		} else {
			fileSize = fmt.Sprintf("(%vb)", size)
		}
	}

	return fileSize
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

