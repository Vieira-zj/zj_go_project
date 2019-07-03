package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// GetBase64Text returns base64 text.
func GetBase64Text(text []byte) string {
	return base64.StdEncoding.EncodeToString(text)
}

// GetURLBase64Text returns base64 text.
func GetURLBase64Text(text []byte) string {
	return base64.URLEncoding.EncodeToString(text)
}

// GetMd5HexText returns md5 hex text.
func GetMd5HexText(text string) string {
	return getMd5EncodedText(text, "hex")
}

// GetBase64MD5Text returns base64 md5 text.
func GetBase64MD5Text(text string) string {
	return getMd5EncodedText(text, "std64")
}

// GetURLBase64MD5Text returns url base64 md5 text.
func GetURLBase64MD5Text(text string) string {
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
