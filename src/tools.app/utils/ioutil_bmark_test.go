package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	myutils "src/tools.app/utils"
)

func BenchmarkReadFile01(b *testing.B) {
	var (
		tmpDirPath = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
		fPath      = filepath.Join(tmpDirPath, "test.out")
	)
	for i := 0; i < b.N; i++ {
		myutils.ReadFileContent(fPath)
	}
}

func BenchmarkReadFile02(b *testing.B) {
	var (
		tmpDirPath = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
		fPath      = filepath.Join(tmpDirPath, "test.out")
	)
	for i := 0; i < b.N; i++ {
		myutils.ReadFileContentBuf(fPath)
	}
}

func BenchmarkReadFile03(b *testing.B) {
	var (
		tmpDirPath = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
		fPath      = filepath.Join(tmpDirPath, "test.out")
	)
	for i := 0; i < b.N; i++ {
		myutils.ReadFileLines(fPath)
	}
}
