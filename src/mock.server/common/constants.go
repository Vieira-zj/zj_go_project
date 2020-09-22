package common

import (
	"encoding/base64"
	"math/rand"
)

const (
	// DataDirPath data dir path of mock.
	DataDirPath = "data"

	// TextContentType http header "Content-Type".
	TextContentType = "Content-Type"
	// TextContentLength http header "Content-Length".
	TextContentLength = "Content-Length"
	// TextContentEncoding http header "Content-Encoding".
	TextContentEncoding = "Content-Encoding"

	// ContentTypeJSON http content type application/json.
	ContentTypeJSON = "application/json; charset=utf-8"
	// ContentTypeTEXT http content type text/plain.
	ContentTypeTEXT = "text/plain; charset=utf-8"
	// ContentTypeHTML http content type text/html.
	ContentTypeHTML = "text/html; charset=uft-8"
	// ContentTypeForm http content type form.
	ContentTypeForm = "application/x-www-form-urlencoded"

	// SvrErrRespMsg error message "Internal Server Error".
	SvrErrRespMsg = "Internal Server Error"
)

// IsProd returns run env is production
func IsProd() bool {
	return RunConfigs.RunEnv == "prod"
}

// CreateMockString returns mock base64 string for size of bytes.
func CreateMockString(size int) string {
	buf := make([]byte, size, size)
	for i := 0; i < size; i++ {
		buf[i] = uint8(rand.Intn(128))
	}
	return getBase64Text(buf)
}

func getBase64Text(text []byte) string {
	return base64.StdEncoding.EncodeToString(text)
}
