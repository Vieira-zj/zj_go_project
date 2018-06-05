package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
)

// GetMd5ForText : get hex md5 for text
func GetMd5ForText(content string) string {
	return getMd5EncodedText(content, "hex")
}

// GetBase64Md5ForText : get base64 md5 for text
func GetBase64Md5ForText(content string) string {
	return getMd5EncodedText(content, "std64")
}

// GetURLBasedMd5ForText : get url based md5 for text
func GetURLBasedMd5ForText(content string) string {
	return getMd5EncodedText(content, "url")
}

func getMd5EncodedText(content string, md5Type string) string {
	md5hash := md5.New()
	md5hash.Write([]byte(content))

	if md5Type == "hex" {
		return hex.EncodeToString(md5hash.Sum(nil))
	} else if md5Type == "std64" {
		return base64.StdEncoding.EncodeToString(md5hash.Sum(nil))
	}
	return base64.URLEncoding.EncodeToString(md5hash.Sum(nil))
}

// GzipEncode : get gzip encode bytes
func GzipEncode(in []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer w.Close()
	_, err := w.Write(in)
	if err != nil {
		return nil, err
	}
	w.Flush()
	return buf.Bytes(), nil
}

// GzipDecode : get gzip decode bytes
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
