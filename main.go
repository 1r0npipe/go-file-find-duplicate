package main

import (
	"fmt"
	"github.com/1r0npipe/go-file-find-duplicate/helper"
)

var (
//delete := flag.Parse("")
)

func main() {

	helper.DuplicatesFind("./", false)
	fmt.Printf("Walked trhough: %d file(-s), found: %d duplicates\n",
		helper.FileCount,
		helper.FilesDuplicates)
}
