package helper

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type fileInfo struct {
	fileSize int64
	fileName string
}

func Finder(filePath string) {
	files := make(map[string]int64)
	err := filepath.Walk(filePath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files[path] = info.Size()
		}
		return nil
	})
	if err != nil {
		fmt.Println("error occurs:", err)
	}
	for name, size := range files {
		fmt.Println(name, size)
	}
}
