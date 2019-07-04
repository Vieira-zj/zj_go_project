package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	myutils "tools.app/utils"
)

func TestReadFileContent(t *testing.T) {
	var (
		tmpDirPath = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
		fPath      = filepath.Join(tmpDirPath, "test.out")
	)

	t.Log("Case01: test read all file content.")
	content, err := myutils.ReadFileContent(fPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(content) == 0 {
		t.Errorf("Failed: file (%s) content is empty!", fPath)
	}
	t.Logf("file (%s) content:\n%s\n", fPath, content)

	t.Log("Case02: test read all file content with buffer.")
	fBytes, err := myutils.ReadFileContentBuf(fPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(fBytes) == 0 {
		t.Errorf("Failed: file (%s) content is empty!", fPath)
	}
	t.Logf("file (%s) content:\n%s\n", fPath, string(fBytes))
}

func TestReadFileLines(t *testing.T) {
	var (
		tmpDirPath = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
		fPath      = filepath.Join(tmpDirPath, "test.out")
	)

	t.Log("Case01: test read all lines from file.")
	lines, err := myutils.ReadFileLines(fPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) == 0 {
		t.Errorf("Failed: file (%s) content is empty!", fPath)
	}
	t.Logf("file (%s) lines count %d, content:\n%v\n", fPath, len(lines), lines)
}
