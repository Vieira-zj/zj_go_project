package main

import (
	"fmt"
	"os"
	"path/filepath"

	myutils "tools.app/utils"
)

func testCreateGzipFile() {
	dirPath := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files/logs")
	dest := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files/logs.tar.gz")

	f, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}
	if err := myutils.CreateGzipFile([]*os.File{f}, dest); err != nil {
		panic(err)
	}
}

func testUngzipFile() {
	// TODO:
}

func main() {
	testCreateGzipFile()

	fmt.Println("utils test DONE.")
}
