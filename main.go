package main

import (
	"github.com/1r0npipe/go-file-find-duplicate/helper"
)

func main() {
	helper.DuplicatesFind("./", false, 1)
}
