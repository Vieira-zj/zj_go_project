package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// GzipEncode get gzip encode bytes.
func GzipEncode(in []byte) ([]byte, error) {
	var outBuf bytes.Buffer
	w := gzip.NewWriter(&outBuf)
	defer w.Close()

	_, err := w.Write(in)
	if err != nil {
		return nil, err
	}
	w.Flush()
	return outBuf.Bytes(), nil
}

// GzipDecode get gzip decode bytes.
func GzipDecode(in []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(in))
	defer r.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		if err != io.ErrUnexpectedEOF {
			return nil, err
		}
	}
	return b, nil
}

// CreateGzipFile gzip compress, and create tar.gz file.
func CreateGzipFile(files []*os.File, dest string) error {
	fOut, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fOut.Close()

	gw := gzip.NewWriter(fOut)
	tw := tar.NewWriter(gw)
	defer func() {
		gw.Close()
		tw.Close()
	}()

	for _, file := range files {
		err := gzipCompress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func gzipCompress(rootFile *os.File, prefix string, tw *tar.Writer) error {
	fInfo, err := rootFile.Stat()
	if err != nil {
		return err
	}
	defer rootFile.Close()

	if fInfo.IsDir() {
		if len(prefix) == 0 {
			prefix = fInfo.Name()
		} else {
			prefix = prefix + "/" + fInfo.Name()
		}
		subFiles, err := rootFile.Readdir(-1)
		if err != nil {
			return err
		}
		for _, sub := range subFiles {
			f, err := os.Open(rootFile.Name() + "/" + sub.Name())
			if err != nil {
				return err
			}
			err = gzipCompress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(fInfo, "")
		if len(prefix) > 0 {
			header.Name = prefix + "/" + header.Name
		}
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, rootFile)
		if err != nil {
			return err
		}
	}
	return nil
}

// UngzipFile uncompress tar.gz file.
func UngzipFile(tarFilePath, dest string) error {
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
		filepath := dest + "/" + hdr.Name
		f, err := createFile(filepath)
		if err != nil {
			return err
		}
		io.Copy(f, tr)
	}
	return nil
}

func createFile(filePath string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(filePath)[0:strings.LastIndex(filePath, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(filePath)
}
