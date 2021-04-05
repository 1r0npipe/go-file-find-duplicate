//Package helper implements the function of removal duplicates of file
//regarding the provided path. It will look into all sub-directories
//two options are available: review duplicates and delete all of them
//no option to delete one by one yet, however you can review first, then delete
package helper

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func ExampleDuplicatesFind() {
	err := DuplicatesFind("./", false, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Walked through: %d file(-s), found: %d duplicates\n",
		FileCount,
		FilesDuplicates)
}

func CreateDuplicates(path, nameDir, nameFile string, dep int) error {
	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("can't change root dir for tests: %v", err)
	}
	for i := 1; i < dep; i++ {
		err := os.Mkdir(nameDir, 0777)
		if err != nil {
			return fmt.Errorf("can't create directory: %v\n", err)
		}
		err = os.Chdir(nameDir)
		if err != nil {
			return fmt.Errorf("can't change directory: %v\n", err)
		}
		_, err = os.Create(nameFile)
		if err != nil {
			return fmt.Errorf("can't create file, because of: %v", err)
		}
	}
	return nil
}
func TestDuplicatesFind(t *testing.T) {
	var path = "/tmp"
	tests := []struct {
		want     int64
		got      int64
		path     string
		nameDir  string
		fileName string
		flag     bool
	}{
		{3, 1, path, "test", "test.txt", false},
		{5, 3, path, "testMore", "testMore.txt", true},
	}

	for _, tt := range tests {
		err := CreateDuplicates(tt.path, tt.nameDir, tt.fileName, int(tt.want))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		err = DuplicatesFind(tt.path, tt.flag, true)
		if err != nil {
			log.Fatalf("Test failed: %v", err)
		}
		if tt.got != tt.want-2 {
			t.Fatalf("Test has failed with want: %d, but got: %d\n",
				tt.want,
				tt.got)
		}
		err = os.Chdir(path)
		if err != nil {
			t.Fatalf("can't change directory: %v", err)
		}
		err = os.RemoveAll(tt.path + string('/') + tt.nameDir)
		if err != nil {
			t.Fatalf("can't remove test directory: %v\n", err)
		}
	}
}
