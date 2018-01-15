package main

import (
	"fmt"
	//"io"
	"os"
	//"path/filepath"
	//"strings"
	"path/filepath"
	"strings"
)

var disabledFound = []string{".git", ".gitignore", ".idea", "README.md"}


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
	err := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		if isDisabled(path) {
			return nil
		}


		fmt.Println(path)

		//fileList = append(fileList, path)
		return nil
	})

	//for _, file := range fileList {
	//	fmt.Println(file)
	//}

	return err
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