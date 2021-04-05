package main

import (
	"flag"
	"github.com/1r0npipe/go-file-find-duplicate/helper"
	"go.uber.org/zap"
	"log"
)

var (
	dirPath = flag.String("dir", "./",
		"Directory to use as root for searching of duplicates, by default current directory")
	del = flag.Bool("delete", false,
		"By default it will just show the duplicates, use \"-delete\" if want to delete")
	debug = flag.Bool("debug", false, "show all duplicated files with information")
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()
	flag.Parse()
	showDegunMsg := false
	if *debug {
		showDegunMsg = true
	}
	err := helper.DuplicatesFind(*dirPath, *del, showDegunMsg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Duplicates searcher statistic",
		zap.Int64("duplicates_found", helper.FilesDuplicates),
		zap.Int64("walked_files", helper.FileCount),
	)
}
