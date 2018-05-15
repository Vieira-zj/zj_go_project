package demos

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	testFilePath = "/Users/zhengjin/Downloads/tmp_files/test.down"
)

// file path handle
func testFilePathHandle() {
	fmt.Println("file name:", filepath.Base(testFilePath))
}

func testCopyBytesToNull() {
	b := bytes.NewReader([]byte("content text put to /dev/null"))
	len, err := io.Copy(ioutil.Discard, b)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Printf("file handler length: %d\n", len)
}

// format and encode
func testURLEncode() {
	// query escape
	baseURL := "http://zj.test.io"
	fmt.Println("path escape:", url.PathEscape(baseURL))
	callbackURL := "http://zj.test.server/callback?k1=v1&k2=v2"
	fmt.Println("query escape:", url.QueryEscape(callbackURL))

	// url with encode query
	query := url.Values{}
	query.Add("mirror", callbackURL)
	fullURL := baseURL + "?" + query.Encode()
	fmt.Println("encode url:", fullURL)
}

// md5 hash, and base64 encode
func getTextMd5Sum(b []byte) string {
	bMd5 := md5.Sum(b)
	fmt.Printf("md5 bytes: %v\n", bMd5)
	return fmt.Sprintf("%x", bMd5)
}

func textEncoded(b []byte, md5Type string) string {
	md5hash := md5.New()
	md5hash.Write(b)
	bMd5 := md5hash.Sum(nil)
	fmt.Printf("md5 bytes: %v\n", bMd5)

	if md5Type == "hex" {
		return hex.EncodeToString(bMd5)
	}
	// Base64编码 使用的字符包括大小写字母各26个, 加上10个数字,
	// 和加号"+", 斜杠"/", 一共64个字符, 等号"="用来作为后缀用途, 其中的 +, /, = 都是需要urlencode的
	if md5Type == "std64" {
		return base64.StdEncoding.EncodeToString(bMd5)
	}
	return base64.URLEncoding.EncodeToString(bMd5)
}

func testMd5Encode() {
	fmt.Println("get text md5:")
	fmt.Println("text md5:", getTextMd5Sum([]byte("ok")))
	fmt.Println("text md5:", textEncoded([]byte("ok"), "hex"))

	fmt.Println("get file content md5:")
	b, _ := ioutil.ReadFile(testFilePath)
	// fmt.Printf("src file data stream bytes: %v\n", b)
	fmt.Println("file md5:", getTextMd5Sum(b)) // byte = 2 hex, 138 = 8a
	fmt.Println("file md5:", textEncoded(b, "hex"))

	fmt.Println("md5 base64 encode:", textEncoded(b, "std64"))
	fmt.Println("md5 base64 url-encode:", textEncoded(b, "url"))
}

// hash check - fnv32
func hashFNV32(text string) uint32 {
	f := fnv.New32()
	f.Write([]byte(text))
	return f.Sum32()
}

func testHashFNV32() {
	url := "www.qiniu.io"
	hashedNum := hashFNV32(url)
	fmt.Printf("fnv32 hash number: %v\n", hashedNum)

	res := hashedNum % 2
	fmt.Printf("mod value: %d\n", res)
}

// get request, read content by range
func getFileByRange(reqURL string) error {
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

func testGetFileByRange() {
	const url = "http://localhost:17890/index1/"
	err := getFileByRange(url)
	if err != nil {
		log.Panicf("error: %v\n", err)
	}
}

// file download and save by http Get
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

	if b, err := ioutil.ReadFile(testFilePath); err == nil {
		fileMd5 := getTextMd5Sum(b)
		fmt.Println("file md5:", fileMd5)
	}
}

// json parser
func testJSONObjToStr() {
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

	if b, err := json.Marshal(group); err == nil {
		fmt.Println("encode string:", string(b))
	}
	if b, err := json.MarshalIndent(group, "", "  "); err == nil {
		fmt.Println("encode string (pretty):", string(b))
	}
}

func testJSONStrToObj1() {
	type Animal struct {
		Name  string `json:"a_name"`
		Order string `json:"a_order"`
	}

	jsonBlob := []byte(`[
		{"a_name": "Platypus", "a_order": "Monotremata"},
		{"a_name": "Quoll",    "a_order": "Dasyuromorphia"}
	]`)
	fmt.Printf("before decode: %s\n", string(jsonBlob))

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("decode object: %+v\n", animals)
	prettyJSON, _ := json.MarshalIndent(animals, "", "  ")
	fmt.Println("pretty json:", string(prettyJSON))

	fmt.Println("animals info:")
	for _, a := range animals {
		fmt.Printf("name=%s, order=%s\n", a.Name, a.Order)
	}
}

func testJSONStrToObj2() {
	type Job struct {
		Title  string   `json:"title"`
		Skills []string `json:"skills"`
	}

	type Person struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Desc Job    `json:"desc"`
	}

	jsonTesters := []byte(`[
		{
		  "id":1, 
		  "name":"person1", 
		  "desc":{
			"title":"tester",
			"skills":["automation test","interface test"]
		  }
		},
		{
		  "id":2, 
		  "name":"person2",
		  "desc":{
			"title":"developer",
			"skills":["python","java","golang"]
		  }
		}
	]`)

	var persons []Person
	err := json.Unmarshal(jsonTesters, &persons)
	if err != nil {
		panic(err.Error())
	}

	for _, p := range persons {
		tmpStr := fmt.Sprintf("%s (id=%d name=%s) skills:", p.Desc.Title, p.ID, p.Name)
		fmt.Println(tmpStr, strings.Join(p.Desc.Skills, ", "))
	}
}

func testJSONStrToRawObj() {
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

	// use interface instead by pre-defined struct, json object map to map[string]interface{}
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

func testArrayStrToSlice() {
	strArray := `["item1", "item2", "item3", "item4", "item5"]`
	// strArray = "null"
	var tmpSlice []string

	json.Unmarshal([]byte(strArray), &tmpSlice)
	fmt.Println("items:")
	if len(tmpSlice) > 0 {
		for idx, item := range tmpSlice {
			fmt.Printf("at %d: %s\n", idx, item)
		}
	} else {
		fmt.Printf("%v\n", tmpSlice)
	}
}

// MainUtils : main for utils
func MainUtils() {
	// testFilePathHandle()
	// testCopyBytesToNull()

	// testURLEncode()
	// testMd5Encode()
	// testHashFNV32()
	// testGetFileByRange()
	// testFileDownload()

	// testJSONObjToStr()
	// testJSONStrToObj1()
	// testJSONStrToObj2()
	// testJSONStrToRawObj()

	// testArrayStrToSlice()

	fmt.Println("utils done.")
}
