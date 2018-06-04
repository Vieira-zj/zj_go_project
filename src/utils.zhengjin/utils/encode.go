package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
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
