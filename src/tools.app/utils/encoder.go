package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// GetMd5ForText get hex md5 for text.
func GetMd5ForText(text string) string {
	return getMd5EncodedText(text, "hex")
}

// GetBase64Md5ForText get base64ed md5 for text.
func GetBase64Md5ForText(text string) string {
	return getMd5EncodedText(text, "std64")
}

// GetURLBase64Md5ForText get url base64ed md5 for text.
func GetURLBase64Md5ForText(text string) string {
	return getMd5EncodedText(text, "url")
}

func getMd5EncodedText(text, md5Type string) string {
	md5hash := md5.New()
	md5hash.Write([]byte(text))
	b := md5hash.Sum(nil)

	if md5Type == "hex" {
		return hex.EncodeToString(b)
	}
	if md5Type == "std64" {
		return base64.StdEncoding.EncodeToString(b)
	}
	return base64.URLEncoding.EncodeToString(b)
}
