package helper

import (
	"fmt"
	"io/fs"
	"os"
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
	FileCount       int64           = 0 // amount of all files to review
	FilesDuplicates int64           = 0 // amount of duplicated files
	WalkedFiles     map[string]File     // list of all files
	Duplicates      = struct {
		sync.RWMutex
		File map[string][]File
	}{File: make(map[string][]File)} // map of duplicated files
)

func DuplicatesFind(filePath string, flag bool, nCPU int) {
	dup := make(chan *File)
	ScanAndFindFiles(filePath)
	var wg sync.WaitGroup
	//go func (ch chan *File) {
	//	ReadDuplicates(ch)
	//}(dup)
	for i := 0; i < nCPU; i++ {
		wg.Add(1)
		go ProcessDuplicates(dup, flag)
		wg.Done()
	}
	ReadDuplicates(dup)
	defer close(dup)
}

// ScanAndFindFiles function is scanning the "filePath" dir
// recursively and "duplicateFiles" provide the output of all duplicates
// to wait while its working, please provide wait group variable
func ScanAndFindFiles(filePath string) {
	var file File
	WalkedFiles = make(map[string]File)
	err := filepath.Walk(filePath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			file.Size = info.Size()
			file.Path = path
			_, filename := filepath.Split(info.Name())
			file.Name = filename
			// file ID will be the string with filename + his size in bytes
			file.Id = filename + strconv.Itoa(int(info.Size()))
			if oneFile, ok := WalkedFiles[file.Id]; ok && (oneFile.Size == info.Size()) {
				Duplicates.File[file.Id] = append(Duplicates.File[file.Id], file)
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

func ReadDuplicates(dupFiles chan *File) {
	for key, value := range Duplicates.File {
		//fmt.Println(value)
		Duplicates.Lock()
		for i := 0; i < len(value); i++ {
			dupFiles <- &value[i]
			//fmt.Println("value",value[i])
		}
		delete(Duplicates.File, key)
		Duplicates.Unlock()
	}
}
func ProcessDuplicates(ch <-chan *File, flag bool) {
	for file := range ch {
		if flag {
			os.Remove(file.Path)
			return
		}
		fmt.Printf("Duplicate: %s, with size %d byte(-s)\n", file.Path, file.Size)
	}

}
