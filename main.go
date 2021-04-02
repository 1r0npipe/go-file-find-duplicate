package main

import (
	"flag"
	"log"
	"github.com/1r0npipe/go-file-find-duplicate/helper"
	"go.uber.org/zap"
)

var (
	dirPath = flag.String("dir", "./",
		"Directory to use as root for searching of duplicates, by default current directory")
	del = flag.Bool("delete", false,
		"By default it will just show the duplicates, use \"-delete\" if want to delete")
)

func main() {
	logger:= zap.NewExample()
	defer logger.Sync()
	
	flag.Parse()
	err := helper.DuplicatesFind(*dirPath, *del)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Duplicates searcher statistic",
		zap.Int64("duplicates_found", helper.FilesDuplicates),
		zap.Int64("walked_files", helper.FileCount),
	)
}
