package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	myutils "tools.app/utils"
)

var (
	tmpDir = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
)

// test, gzip compress and decompress
func testGzipCompressFile() {
	src := filepath.Join(tmpDir, "pics/upload.jpg")
	dest := filepath.Join(tmpDir, "pics/upload.tar.gz")

	f, err := os.Open(src)
	if err != nil {
		fmt.Println("read src file error:", err.Error())
	}
	err = myutils.CreateGzipFile([]*os.File{f}, dest)
	if err != nil {
		fmt.Println("comporess error:", err.Error())
	}
}

func testGzipCompressDir() {
	src := filepath.Join(tmpDir, "tmp_dir")
	dest := filepath.Join(tmpDir, "tmp_dir.tar.gz")

	if f, err := os.Open(src); err == nil {
		fmt.Printf("compress file (%s) with tar.gz\n", src)
		err := myutils.CreateGzipFile([]*os.File{f}, dest)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func testGzipDecompress() {
	src := filepath.Join(tmpDir, "tmp_dir.tar.gz")
	dest := tmpDir

	err := myutils.UngzipFile(src, dest)
	if err != nil {
		panic(err)
	}
	fmt.Println("decompress to:", dest)
}

// test, gzip encode and decode
func testGzipCoder() {
	fPath := filepath.Join(tmpDir, "test.out")

	// srcb := []byte("gzip encode and decode, test.")
	srcb, err := ioutil.ReadFile(fPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("src length:", len(srcb))

	destb, err := myutils.GzipEncode(srcb)
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip encode length:", len(destb))

	b, err := myutils.GzipDecode(destb)
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip decode length:", len(b))
	if len(b) <= 128 {
		fmt.Println("encode and decode bytes:", string(b))
	}
}

func main() {

	// testGzipCompressFile()
	// testGzipCompressFile()
	// testGzipDecompress()

	// testGzipCoder()

	fmt.Println("golang utils test DONE.")
}