package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// GzipEncode returns gzip encode bytes.
func GzipEncode(in []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer w.Close()
	w.Comment = "gzip encode for compress text."

	if _, err := w.Write(in); err != nil {
		return nil, err
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GzipDecode returns gzip decode bytes.
func GzipDecode(in []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(in))
	defer r.Close()
	if err != nil {
		return nil, err
	}
	if len(r.Comment) > 0 {
		fmt.Println("Comment:", r.Comment)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		if err != io.ErrUnexpectedEOF {
			return nil, err
		}
	}
	return b, nil
}

// CompressGzipFile gzip compress, and create tar.gz file.
func CompressGzipFile(files []*os.File, dest string) error {
	newFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer newFile.Close()

	gw := gzip.NewWriter(newFile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		err = compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	fnGetPrefix := func(prefix, text string) string {
		if len(prefix) > 0 {
			return prefix + "/" + text
		}
		return text
	}

	fInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fInfo.IsDir() {
		prefix = fnGetPrefix(prefix, fInfo.Name())
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(fInfo, "")
		if err != nil {
			return err
		}
		header.Name = fnGetPrefix(prefix, header.Name)
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeCompressGzipFile de-compress tar.gz file.
func DeCompressGzipFile(tarFilePath, dest string) error {
	tarFile, err := os.Open(tarFilePath)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	gr, err := gzip.NewReader(tarFile)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filepath := filepath.Join(dest, hdr.Name)
		f, err := createFile(filepath)
		if err != nil {
			return err
		}
		io.Copy(f, tr)
	}
	return nil
}

func createFile(filePath string) (*os.File, error) {
	parentAbsPath := string([]rune(filePath)[0:strings.LastIndex(filePath, "/")])
	err := os.MkdirAll(parentAbsPath, 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(filePath)
}
