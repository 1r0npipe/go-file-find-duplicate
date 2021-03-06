package helper

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
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

// DuplicatesFind main function to find all duplicates
// Input: "filePath" directory to start searching of duplicates
// "flag" - if true remove, if false - show
func DuplicatesFind(fileSystem fs.FS, flag bool) error {
	dup := make(chan *File)
	err := ScanAndFindFiles(fileSystem)
	if err != nil {
		return fmt.Errorf("can't search due to this error %v", err)
	}
	go ReadDuplicates(dup)

	for file := range dup {
		err := ProcessDuplicates(file, flag)
		if err != nil {
			return fmt.Errorf("issue with processing of duplicates: %v", err)
		}
	}
	return nil
}

// ScanAndFindFiles function is scanning the "filePath" dir
// recursively and "duplicateFiles" provide the output of all duplicates
// to wait while its working, please provide wait group variable
//func ScanAndFindFiles(filePath string) error {
func ScanAndFindFiles(fileSystem fs.FS) error {
	var file File
	WalkedFiles = make(map[string]File)
	//err := filepath.Walk(filePath, func(path string, info fs.FileInfo, err error) error {
	err := fs.WalkDir(fileSystem, ".", func(path string, info fs.DirEntry, err error) error {
		fileInfo, _ := info.Info()
		if !info.IsDir() {
			file.Size = fileInfo.Size()
			file.Path = path
			_, filename := filepath.Split(info.Name())
			file.Name = filename
			// file ID will be the string with filename + his size in bytes
			file.Id = filename + strconv.Itoa(int(fileInfo.Size()))
			if oneFile, ok := WalkedFiles[file.Id]; ok && (oneFile.Size == fileInfo.Size()) {
				Duplicates.File[file.Id] = append(Duplicates.File[file.Id], file)
				atomic.AddInt64(&FilesDuplicates, 1)
			}
			WalkedFiles[file.Id] = file
		}
		atomic.AddInt64(&FileCount, 1)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error occurs: %v", err)
	}
	return nil
}

// ReadDuplicates read all duplicates from structure
// and push it to channel - "dupFiles"
func ReadDuplicates(dupFiles chan *File) {
	for key, value := range Duplicates.File {
		Duplicates.Lock()
		for i := 0; i < len(value); i++ {
			dupFiles <- &value[i]
		}
		delete(Duplicates.File, key)
		Duplicates.Unlock()
	}
	defer close(dupFiles)
}

// ProcessDuplicates process with duplicates:
// if flag is true - delete, otherwise just show
// also it needs channel to read the duplicates from
func ProcessDuplicates(file *File, flag bool) error {
	if flag {
		err := os.Remove(file.Path)
		if err != nil {
			return fmt.Errorf("can't remove file: %s by this error: %v",
				file.Path, err)
		}
		return nil
	}
	fmt.Printf("Duplicate: %s, with size %d byte(-s);\n",
		file.Path,
		file.Size)
	return nil
}
