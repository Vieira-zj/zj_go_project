package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	myutils "tools.app/utils"
)

func TestReadFileContent(t *testing.T) {
	t.Skip("skip TestReadFileContent.")
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
	t.Skip("skip TestReadFileLines.")
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

func TestWriteToNewFile(t *testing.T) {
	t.Skip("skip TestWriteToNewFile.")
	var (
		tmpDirPath = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
		fPath      = filepath.Join(tmpDirPath, "write_test.out")
	)

	t.Log("Case01: write content to a file with overwrite=false.")
	content := "1|one\n2|two\n3|three\n4|four\n5|five"
	err := myutils.WriteContentToFile(fPath, content, false)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Case02: write content to an exist file with overwrite=false.")
	err = myutils.WriteContentToFile(fPath, content, false)
	if err == nil {
		t.Error("Failed: write to an exist file!")
	}
	t.Logf("Error: %v\n", err)

	t.Log("Clearup files")
	err = myutils.DeleteFile(fPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAppendToFile(t *testing.T) {
	t.Skip("skip TestAppendToFile.")
	var (
		tmpDirPath = filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
		fPath      = filepath.Join(tmpDirPath, "append_test.out")
	)

	t.Log("Case01: append content to a new file.")
	content1 := "1|one\n2|two\n3|three\n4|four\n5|five"
	n, err := myutils.AppendContentToFile(fPath, content1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("write bytes count: %d\n", n)

	t.Log("Case02: append content to an exist file.")
	content2 := "\n6|six\n7|seven\n8|eight"
	n, err = myutils.AppendContentToFile(fPath, content2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("append bytes count: %d\n", n)

	t.Log("Step: check file content.")
	rText, err := myutils.ReadFileContent(fPath)
	if err != nil {
		t.Fatal(err)
	}
	if rText != (content1 + content2) {
		t.Log("Failed: write and append content to file.")
	}
	t.Logf("output file content:\n%s\n", rText)

	t.Log("Clearup files")
	err = myutils.DeleteFile(fPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCurPath(t *testing.T) {
	t.Log("Case01: test get current run abs path.")
	t.Log("current path:", myutils.GetCurPath())
}
