package mocks

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// GetContentMd5 : get md5 value from input
func GetContentMd5(input string, md5Type string) string {
	md5hash := md5.New()
	md5hash.Write([]byte(input))

	if md5Type == "hex" {
		return hex.EncodeToString(md5hash.Sum(nil))
	} else if md5Type == "std64" {
		return base64.StdEncoding.EncodeToString(md5hash.Sum(nil))
	}
	return base64.URLEncoding.EncodeToString(md5hash.Sum(nil))
}
