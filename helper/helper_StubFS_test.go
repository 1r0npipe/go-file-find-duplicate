// +build integration

package helper

import (
	"io/fs"
	"log"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"testing"
)
// Defined duplicated below:
// Duplicate: test1/clone2, with size 9 byte(-s);
// Duplicate: test3/clone2, with size 9 byte(-s);
// Duplicate: test2/unique, with size 9 byte(-s);
// Duplicate: test3/unique, with size 9 byte(-s);
// Duplicate: test3/clone1, with size 9 byte(-s);
var fakeFS = fstest.MapFS{
	// root-folder
	"testDir": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	"clone2": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	//test1
	"test1": {
		Mode: fs.ModeDir,
	},
	"test1/clone1": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	"test1/clone2": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	"test1/unique": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	//test2
	"test2": {
		Mode: fs.ModeDir,
	},
	"test2/clone1": {
		Data: []byte("clone1"),
		Mode: fs.ModeIrregular,
	},
	"test2/clone3": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	"test2/unique": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	//test3
	"test3": {
		Mode: fs.ModeDir,
	},
	"test3/clone1": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	"test3/clone2": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	"test3/unique": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
	"test3/unique2": {
		Data: []byte("Text-test"),
		Mode: fs.ModeIrregular,
	},
}

func TestFindDupByFakeFS(t *testing.T) {
	testsTab := []struct {
		got  int64
	}{
		{got: 5},
	}
	for _, tt := range testsTab {
		err := DuplicatesFind(fakeFS, false)
		if err != nil {
			log.Fatal("Can't get duplicates")
		}
		assert.Equal(t, tt.got, FilesDuplicates, "Not same amount of found duplicates")
	}
}
