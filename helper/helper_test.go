//Package helper implements the function of removal duplicates of file
//regarding the provided path. It will look into all sub-directories
//two options are available: review duplicates and delete all of them
//no option to delete one by one yet, however you can review first, then delete
package helper

import "testing"

func ExampleDuplicatesFind() {

}
func TestDuplicatesFind(t *testing.T) {
	type args struct {
		filePath string
		flag     bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestProcessDuplicates(t *testing.T) {
	type args struct {
		file *File
		flag bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestReadDuplicates(t *testing.T) {
	type args struct {
		dupFiles chan *File
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestScanAndFindFiles(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
