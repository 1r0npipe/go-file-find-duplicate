package main

import (
	"flag"
	"fmt"
	"github.com/1r0npipe/go-file-find-duplicate/helper"
	"log"
	"os"
)

var (
	dirPath = flag.String("dir", "./",
		"Directory to use as root for searching of duplicates, by default current directory")
	del = flag.Bool("delete", false,
		"By default it will just show the duplicates, use \"-delete\" if want to delete")
)

func main() {
	flag.Parse()
	fileSystem := os.DirFS(*dirPath)
	err := helper.DuplicatesFind(fileSystem, *del)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Walked through: %d file(-s), found: %d duplicates\n",
		helper.FileCount,
		helper.FilesDuplicates)
}
