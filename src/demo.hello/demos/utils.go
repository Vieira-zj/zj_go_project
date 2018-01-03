package demos

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	testFilePath = "/Users/zhengjin/Downloads/tmp_files/test.down"
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
	fileMd5, _ := getFileMd5(testFilePath)
	fmt.Println("file md5:", fileMd5)

	b, _ := ioutil.ReadFile(testFilePath)
	fmt.Println("hex encoded md5:", getEncodedMd5(b, "hex"))
}

// file download
func fileDownloadAndSave(reqURL, filePath string) error {
	fmt.Printf("request url: %s\n", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return err
	}
	fmt.Printf("ret code: %d\n", resp.StatusCode)

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Printf("saving at: %s\n", filePath)
	io.Copy(f, resp.Body)
	defer resp.Body.Close()

	fmt.Println("downfile file done.")
	return nil
}

func testFileDownload() {
	query := &url.Values{}
	query.Add("uid", "1380469261")
	query.Add("bucket", "publicbucket_z0")
	query.Add("url", "http://10.200.20.21:17890/index4/")
	url := "http://qiniuproxy.kodo.zhengjin.cs-spock.cloudappl.com/mirror?"
	url += query.Encode()

	if err := fileDownloadAndSave(url, testFilePath); err != nil {
		panic(err.Error())
	}

	fileMd5, _ := getFileMd5(testFilePath)
	fmt.Println("file md5:", fileMd5)
}

// MainUtils : main for utils
func MainUtils() {
	// testMd5Check()
	testFileDownload()

	fmt.Println("utils done.")
}
