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
	err := DuplicatesFind("./", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Walked trhough: %d file(-s), found: %d duplicates\n",
		FileCount,
		FilesDuplicates)
}

func CreateDuplicates(path, nameDir, nameFile string, dep int) error {
	for i := 1; i < dep; i++ {
		err := os.Mkdir(path+nameDir, 0777)
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
	//path, err := os.Getwd()
	var path = "./"
	tests := []struct {
		want     int64
		path     string
		nameDir  string
		fileName string
		flag     bool
	}{
		{3, path, "test", "test.txt", false},
		{5, path, "testMore", "testMore.txt", true},
	}

	for _, tt := range tests {
		err := CreateDuplicates(tt.path, tt.nameDir, tt.fileName, int(tt.want))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		err = DuplicatesFind("./", tt.flag)
		if err != nil {
			log.Fatalf("Test failed: %v", err)
		}
		if FilesDuplicates != tt.want {
			t.Fatalf("Test has failed with want: %d, but got: %d\n",
				tt.want,
				FilesDuplicates)
		}
	}
}
