package demos

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

// md5 check
func getFileMd5(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(b)), nil
}

func getEncodedMd5(b []byte, md5Type string) string {
	md5hash := md5.New()
	md5hash.Write(b)
	bMd5 := md5hash.Sum(nil)

	if md5Type == "hex" {
		return hex.EncodeToString(bMd5)
	}
	if md5Type == "std64" {
		return base64.StdEncoding.EncodeToString(bMd5)
	}
	return base64.URLEncoding.EncodeToString(bMd5)
}

func testMd5Check() {
	path := "/Users/zhengjin/Downloads/tmp_files/test.mp3"
	fileMd5, _ := getFileMd5(path)
	fmt.Println("file md5:", fileMd5)

	b, _ := ioutil.ReadFile(path)
	fmt.Println("hex encoded md5:", getEncodedMd5(b, "hex"))
}

// MainUtils : main for utils
func MainUtils() {
	testMd5Check()

	fmt.Printf("utils done.")
}
