package helper

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func finder(filePath sting) {
	var files []string
	err := filepath.Walk(filePath, func(path string, info fs.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println("error occurs:", err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
