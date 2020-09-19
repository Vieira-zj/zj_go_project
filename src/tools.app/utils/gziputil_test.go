package utils_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	myutils "src/tools.app/utils"
)

// TestGzipCoder, gzip encode and decode test.
func TestGzipCoder(t *testing.T) {
	b := []byte("gzip encode and decode test.")
	t.Logf("source bytes size: %d\n", len(b))

	t.Log("Case01: test gzip encode bytes.")
	encodeB, err := myutils.GzipEncode(b)
	if err != nil {
		t.Fatalf("Failed gzip encode: %v\n", err)
	}
	t.Logf("gzip encode bytes size: %d\n", len(encodeB))

	t.Log("Case02: test gzip decode bytes.")
	decodeB, err := myutils.GzipDecode(encodeB)
	if err != nil {
		t.Fatalf("Failed gzip decode: %v\n", err)
	}
	t.Logf("gzip decode bytes size: %d\n", len(decodeB))

	if len(b) != len(decodeB) {
		t.Errorf("decode bytes size not matched: src=%d,decode=%d\n", len(b), len(decodeB))
	}
	if string(b) != string(decodeB) {
		t.Error("decode text not matched.")
	}
	t.Logf("decode text: %s\n", string(decodeB))
}

// TestCompressGzipFile gzip compress and decompress test for single file.
func TestCompressGzipFile(t *testing.T) {
	t.Log("Case01: test gzip compress a single file.")
	tmpDir := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
	comSrc := filepath.Join(tmpDir, "gaokeyong.jmx")
	comDest := filepath.Join(tmpDir, "gaokeyong.tar.gz")

	t.Log("Step01: check source file.")
	f, err := os.Open(comSrc)
	if err != nil {
		t.Fatalf("Failed open src file: %v\n", err)
	}
	info, err := f.Stat()
	if err != nil {
		t.Fatalf("Failed stat src file: %v\n", err)
	}
	srcName := info.Name()
	srcSize := info.Size()
	t.Logf("src file: name=%s, size=%d\n", srcName, srcSize)

	t.Log("Step02: run gzip compress.")
	err = myutils.CompressGzipFile([]*os.File{f}, comDest)
	if err != nil {
		t.Fatalf("Failed gzip compress: %v\n", err)
	}

	t.Log("Step03: check .gar.gz file.")
	gzipFile, err := os.Open(comDest)
	if err != nil {
		t.Fatalf("Failed open gzip .tar.gz file: %v\n", err)
	}
	if fInfo, err := gzipFile.Stat(); err == nil {
		t.Logf("gzip file name: %s\n", fInfo.Name())
	}
	time.Sleep(time.Second)

	t.Log("Case02: test gzip decompress a .tar.gz file (contains a single file).")
	decomDir := filepath.Join(tmpDir, "tmp_results")
	decomDest := filepath.Join(decomDir, srcName)

	t.Log("Step01: run gzip decompress.")
	err = myutils.DeCompressGzipFile(comDest, decomDir)
	if err != nil {
		t.Fatalf("Failed gzip decomporess: %v\n", err)
	}

	t.Log("Step02: check decompress file.")
	ungzipFile, err := os.Open(decomDest)
	if err != nil {
		t.Fatalf("Failed open ungzip file: %v\n", err)
	}
	info, err = ungzipFile.Stat()
	if err != nil {
		t.Fatalf("Failed stat ungzip file: %v\n", err)
	}

	if info.Size() != srcSize {
		t.Error("gzip Compress and Decompress file size is not matched!")
	}
	t.Logf("decompress file size: %d\n", info.Size())
}

// TestGzipCompressDir gzip compress and decompress test for directory (inculde files).
func TestGzipCompressDir(t *testing.T) {
	t.Log("Case01: test gzip compress multiple files in a dir.")
	tmpDir := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
	comSrc := filepath.Join(tmpDir, "logs")
	comDest := filepath.Join(tmpDir, "logs.tar.gz")

	t.Log("Step01: check source dir (contains multiple files)")
	f, err := os.Open(comSrc)
	if err != nil {
		t.Fatalf("Failed open src file: %v\n", err)
	}

	t.Log("Step02: run gzip compress.")
	err = myutils.CompressGzipFile([]*os.File{f}, comDest)
	if err != nil {
		t.Fatalf("Failed gzip compress: %v\n", err)
	}

	t.Log("Step03: check compress .tar.gz file.")
	gzipFile, err := os.Open(comDest)
	if err != nil {
		t.Fatalf("Failed open .tar.gz file: %v\n", err)
	}
	if fInfo, err := gzipFile.Stat(); err == nil {
		t.Logf("gzip file name: %s\n", fInfo.Name())
	}

	t.Log("Case02: test gzip decompress a .tar.gz file (contains multiple files).")
	decomDir := filepath.Join(tmpDir, "tmp_results")

	t.Log("Step01: run gzip decompress.")
	if err := myutils.DeCompressGzipFile(comDest, decomDir); err != nil {
		t.Fatalf("Failed gzip decomporess: %v\n", err)
	}
}
