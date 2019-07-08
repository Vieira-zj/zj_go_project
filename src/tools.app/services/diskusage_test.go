package services_test

import (
	"os"
	"path/filepath"
	"testing"

	mysvc "tools.app/services"
)

func TestListFiles(t *testing.T) {
	t.Logf("Case01: test list files for a directory.")
	path := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
	diskUage := mysvc.NewDiskUsage()

	files, err := diskUage.ListFiles(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Errorf("no files found in dir (%s)", path)
	}

	t.Logf("files in dir (%s):", path)
	for _, f := range files {
		t.Log(f.Name())
	}
}

func TestPrintDirDiskUsage(t *testing.T) {
	t.Logf("Case01: print disk space usage for a directory.")
	path := filepath.Join(os.Getenv("HOME"), "Downloads")
	// path := os.Getenv("HOME")

	// files count verify cmd:
	// find ~/Downloads -type f | wc -l
	diskUage := mysvc.NewDiskUsage()
	if err := diskUage.PrintDirDiskUsage(path); err != nil {
		t.Fatal(err)
	}
}
