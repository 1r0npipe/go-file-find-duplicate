package helper

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"sync"
)

type File struct {
	Size int64
	Path string
	Name string
	Id   string
}

var (
	//duplicates = struct {
	//	//sync.RWMutex
	//	m map[string][]string
	//}{m : make(map[string][]string)}
	FileCount       int64           = 0 // amount of all files to review
	FilesDuplicates int64           = 0 // amount of duplicated files
	WalkedFiles     map[string]File     // list of all files
	Duplicates      = struct {
		sync.RWMutex
		File map[string][]File
	}{File: make(map[string][]File)} // map of duplicated files
)

func DuplicatesFind(filePath string, flag bool, nCPU int) {
	var wg sync.WaitGroup
	dup := make(chan File)
	ScanAndFindFiles(filePath, dup, &wg)
	wg.Wait()
	if flag {
		for dupFile := range dup {
			fmt.Println(dupFile)
		}
		return
	}
	for i := 0; i < nCPU; i++ {
		go ProcessDuplicates(dup)
	}

}

// ScanAndFindFiles function is scanning the "filePath" dir
// recursively and "duplicateFiles" provide the output of all duplicates
// to wait while its working, please provide wait group variable
func ScanAndFindFiles(filePath string, duplicateFiles chan File, wg *sync.WaitGroup) {
	var file File
	wg.Add(1)
	defer func() {
		close(duplicateFiles)
		wg.Done()
	}()
	WalkedFiles = make(map[string]File)
	err := filepath.Walk(filePath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			file.Size = info.Size()
			file.Path = path
			_, filename := filepath.Split(info.Name())
			// file ID will be the string with filename + his size in bytes
			file.Id = filename + strconv.Itoa(int(info.Size()))
			if oneFile, ok := WalkedFiles[file.Id]; ok && (oneFile.Size == info.Size()) {
				Duplicates.File[file.Id] = append(Duplicates.File[file.Id], file)
				duplicateFiles <- file
				FilesDuplicates++
			}
			WalkedFiles[file.Id] = file
		}
		FileCount++
		return nil
	})
	if err != nil {
		fmt.Println("error occurs:", err)
	}
}

func ProcessDuplicates(ch <-chan File) {
	for key, value := range Duplicates.File {
		Duplicates.Lock()
		dubFile := <-ch
		fmt.Println(key, value, dubFile)
		Duplicates.Unlock()
	}
}
