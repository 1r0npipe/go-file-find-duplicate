package main

import (
	"github.com/1r0npipe/go-file-find-duplicate/helper"
)

func main() {
	helper.DuplicatesFind("./", false, 1)
	//fmt.Println(helper.FileCount)
	//for key, _ := range helper.Duplicates.File {
	//	fmt.Println(helper.Duplicates.File[key])
	//}

}
