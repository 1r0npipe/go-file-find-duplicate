package main

import (
	"fmt"
	"github.com/1r0npipe/go-file-find-duplicate/helper"
	"log"
)

var (
//delete := flag.Parse("")
)

func main() {

	err := helper.DuplicatesFind("./", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Walked trhough: %d file(-s), found: %d duplicates\n",
		helper.FileCount,
		helper.FilesDuplicates)
}
