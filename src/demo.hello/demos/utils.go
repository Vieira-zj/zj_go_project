package demos

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
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
	bMd5 := md5hash.Sum(nil) // bin

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
	fmt.Println("file md5:", getEncodedMd5(b, "hex"))
	fmt.Println("md5 base64 encode:", getEncodedMd5(b, "std64"))
}

// get request, read content by range
func getContentByRange(reqURL string) error {
	log.Printf("request url: %s\n", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return err
	}
	log.Printf("ret code: %d\n", resp.StatusCode)
	defer resp.Body.Close()

	var (
		rRange int64 = 64
		total  int64
	)
	for i := 0; i < 100; i++ {
		log.Println("read and wait...")
		time.Sleep(time.Duration(500) * time.Millisecond)
		length, err := io.CopyN(os.Stdout, resp.Body, rRange)
		fmt.Println()

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		total += length
	}

	log.Printf("content length: %v\n", total)
	return nil
}

func testGetContentByRange() {
	const url = "http://localhost:17890/index1/"
	err := getContentByRange(url)
	if err != nil {
		log.Panicf("error: %v\n", err)
	}
}

// file download and save
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

// json parser
func testJSONObjectToString() {
	type ColorGroup struct {
		ID     int      `json:"cg_id" bson:"cg_id"`
		Name   string   `json:"cg_name" bson:"cg_name"`
		Colors []string `json:"cg_colors" bson:"cg_colors"`
	}

	group := &ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	fmt.Printf("before encode: %+v\n", group)

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("encode string: %s\n", string(b))
}

func testJSONStringToObject() {
	jsonBlob := []byte(`[
		{"a_name": "Platypus", "a_order": "Monotremata"},
		{"a_name": "Quoll",    "a_order": "Dasyuromorphia"}
	]`)
	fmt.Printf("before decode: %s\n", string(jsonBlob))

	type Animal struct {
		Name  string `json:"a_name"`
		Order string `json:"a_order"`
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("decode object: %+v\n", animals)

	fmt.Println("animals info:")
	for _, a := range animals {
		fmt.Printf("name=%s, order=%s\n", a.Name, a.Order)
	}
}

func testJSONStringToRawObject() {
	type skill struct {
		Name  string `json:"skill_name"`
		Level string `json:"skill_level"`
	}

	type tester struct {
		ID     string  `json:"tester_id"`
		Name   string  `json:"tester_name"`
		Skills []skill `json:"tester_skills"`
	}

	t := tester{
		ID:   "id01",
		Name: "tester01",
		Skills: []skill{
			skill{
				Name:  "automation",
				Level: "junior",
			},
			skill{
				Name:  "manual",
				Level: "senior",
			},
		},
	}

	b, err := json.Marshal(t)
	if err != nil {
		log.Panicf("error: %v\n", err)
		return
	}
	fmt.Printf("json string: %s\n", string(b))

	// use interface instead by struct, json object map to map[string]interface{}
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Panicf("panic: %v\n", err)
	}
	fmt.Printf("json object: %v\n", m)

	testers := m["tester_id"]
	fmt.Printf("stills for %s:\n", testers.(string))
	skills := m["tester_skills"]
	for idx, skill := range skills.([]interface{}) {
		name := skill.(map[string]interface{})["skill_name"]
		fmt.Printf("%d) %s\n", idx, name.(string))
	}
}

// MainUtils : main for utils
func MainUtils() {
	// testMd5Check()
	// testGetContentByRange()
	// testFileDownload()

	// testJSONObjectToString()
	// testJSONStringToObject()
	// testJSONStringToRawObject()

	fmt.Println("utils done.")
}
