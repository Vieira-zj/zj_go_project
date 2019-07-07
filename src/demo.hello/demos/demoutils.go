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
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/larspensjo/config"
)

// file handlers
func testGetBaseFileName() {
	testFilePath := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files/test.out")
	if _, err := os.Stat(testFilePath); os.IsNotExist(err) {
		fmt.Println("file not found:", testFilePath)
	} else {
		fmt.Println("file name:", filepath.Base(testFilePath))
	}
}

// write text to /dev/null
func testCopyBytesToNull() {
	b := bytes.NewReader([]byte("content put to /dev/null"))
	n, err := io.Copy(ioutil.Discard, b)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Printf("\ncopied bytes size: %d\n", n)
}

// url format and encode
func testURLParser() {
	// escape
	baseURL := "http://zj.test.io"
	fmt.Println("\nescaped url:", url.PathEscape(baseURL))
	callbackURL := "http://zj.test.server/callback?k1=v1&k2=v2"
	fmt.Println("escaped query:", url.QueryEscape(callbackURL))

	// encode query
	query := url.Values{}
	query.Add("mirror", callbackURL)
	fullURL := baseURL + "?" + query.Encode()
	fmt.Println("\nurl with encoded query:", fullURL)
}

// md5 hash, and base64 encode
func testMD5Encode() {
	// hex
	fmt.Println("\nhex value:")
	fmt.Printf("hex: %x\n", "ok")
	fmt.Println("hex:", hex.EncodeToString([]byte("ok")))

	// md5 (byte = 2 hex, 138 = 8a)
	const tmpstr = "hello world"
	fmt.Println("\nmd5 encode:")
	fmt.Println("#1 md5 hex:", getTextMD5Sum([]byte(tmpstr)))
	fmt.Println("#2 md5 hex:", textEncode([]byte(tmpstr), "hex"))

	// base64
	fmt.Println("\nget base64 and url-base64 encode:")
	fmt.Println("base64 encode:", textEncode([]byte(tmpstr), "std64"))
	fmt.Println("url-base64 encode:", textEncode([]byte(tmpstr), "url"))
}

func textEncode(b []byte, encodeType string) string {
	md5hash := md5.New()
	md5hash.Write(b)
	bMd5 := md5hash.Sum(nil)
	fmt.Printf("md5 bytes: %v\n", bMd5)

	if encodeType == "std64" {
		return base64.StdEncoding.EncodeToString(bMd5)
	}
	if encodeType == "url" {
		return base64.URLEncoding.EncodeToString(bMd5)
	}
	return hex.EncodeToString(bMd5)
}

func getTextMD5Sum(b []byte) string {
	bMd5 := md5.Sum(b)
	fmt.Printf("md5 bytes: %v\n", bMd5)
	return fmt.Sprintf("%x", bMd5)
}

// hash check: fnv32
func testHashFNV32() {
	url := "www.qiniu.io"
	hashedNum := hashFNV32(url)
	fmt.Printf("\nfnv32 hash results: %v\n", hashedNum)
	fmt.Printf("fnv32 hash with mod: %d\n", hashedNum%2)
}

func hashFNV32(text string) uint32 {
	f := fnv.New32()
	f.Write([]byte(text))
	return f.Sum32()
}

// reg expression
func testRegExp() {
	const testStr = "test1, hello, test2,test3, 99,test4"
	if r, err := regexp.Compile(`hello|world`); err == nil {
		fmt.Println("\n#1 string matched:", r.MatchString(testStr))
	}

	if r, err := regexp.Compile(`(\d\d)`); err == nil {
		fmt.Println("#2 number found:", r.FindString(testStr))
	}
}

// get request, read resp content by range
func testMockGetRespBytesByRange() {
	const url = "http://localhost:17891/index1/"
	err := readRespBytesByRange(url)
	if err != nil {
		log.Panicf("error: %v\n", err)
	}
}

func readRespBytesByRange(reqURL string) error {
	log.Printf("\nrequest url: %s\n", reqURL)
	resp, err := http.Get(reqURL)
	log.Printf("ret code: %d\n", resp.StatusCode)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	var (
		rRange int64 = 64
		total  int64
	)
	for i := 0; i < 100; i++ {
		log.Println("read and wait...")
		len, err := io.CopyN(os.Stdout, resp.Body, rRange)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		total += len
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

	log.Printf("resp content length: %v\n", total)
	return nil
}

// file download and save by http Get
func testMockFileDownload() {
	testFilePath := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files/test.out")

	query := &url.Values{}
	query.Add("uid", "1380469261")
	query.Add("bucket", "publicbucket_z0")
	query.Add("url", "http://10.200.20.21:17891/index4/") // 回源地址
	url := "http://qiniuproxy.kodo.zhengjin.cs-spock.cloudappl.com/mirror?"
	url += query.Encode()
	if err := fileDownloadAndSave(url, testFilePath); err != nil {
		panic(err)
	}

	if b, err := ioutil.ReadFile(testFilePath); err == nil {
		fmt.Println("download file md5:", getTextMD5Sum(b))
	}
}

func fileDownloadAndSave(reqURL, filePath string) error {
	fmt.Printf("\ndownload url: %s\n", reqURL)
	resp, err := http.Get(reqURL)
	fmt.Printf("ret code: %d\n", resp.StatusCode)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	saveFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer saveFile.Close()

	n, err := io.Copy(saveFile, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("copy resp %d bytes to file: %s\n", n, filePath)
	return nil
}

// json parser
func testJSONMarshal() {
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
	fmt.Printf("\njson object: %+v\n", *group)

	if b, err := json.Marshal(group); err == nil {
		fmt.Println("marshal json string:", string(b))
	}
	if b, err := json.MarshalIndent(group, "", "  "); err == nil {
		fmt.Println("marchal josn string (pretty):", string(b))
	}
}

func testJSONUnmarshal01() {
	type Animal struct {
		Name  string `json:"a_name"`
		Order string `json:"a_order"`
	}

	jsonBlob := []byte(`[
		{"a_name": "Platypus", "a_order": "Monotremata"},
		{"a_name": "Quoll",    "a_order": "Dasyuromorphia"}
	]`)
	fmt.Printf("\njson string: %s\n", string(jsonBlob))

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nunmarshal json object: %+v\n", animals)

	fmt.Println("animals info:")
	for _, a := range animals {
		fmt.Printf("name=%s, order=%s\n", a.Name, a.Order)
	}

	prettyJSON, _ := json.MarshalIndent(animals, "", "  ")
	fmt.Println("\nmarchal pretty json:", string(prettyJSON))
}

func testJSONUnmarshal02() {
	type Job struct {
		Title  string   `json:"title"`
		Skills []string `json:"skills"`
	}

	type Person struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Desc Job    `json:"desc"`
	}

	jsonBlob := []byte(`[
		{
		  "id":1, 
		  "name":"person1", 
		  "desc":{
			"title":"tester",
			"skills":["web test","interface test"]
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
	err := json.Unmarshal(jsonBlob, &persons)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\nperson info:")
	for _, p := range persons {
		tmp := fmt.Sprintf("id=%d name=%s (%s), and skills:", p.ID, p.Name, p.Desc.Title)
		fmt.Println(tmp, strings.Join(p.Desc.Skills, ","))
	}
}

func testJSONUnmarshalToSlice() {
	strJSONArray := `["item1", "item2", "item3", "item4", "item5"]`
	// strJSONArray = "null"
	var tmpSlice []string
	if err := json.Unmarshal([]byte(strJSONArray), &tmpSlice); err != nil {
		panic(err)
	}

	fmt.Println("\nitems:")
	if len(tmpSlice) > 0 {
		for _, item := range tmpSlice {
			fmt.Println(item)
		}
	} else {
		fmt.Println("items is empty!")
	}
}

func testJSONUnmarshalToRawObj() {
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

	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Panicf("error: %v\n", err)
		return
	}
	fmt.Printf("\nmarshal string:\n%s\n", string(b))

	// unmarshal by interface instead pre-defined struct
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Panicf("panic: %v\n", err)
	}
	fmt.Printf("\nunmarshal raw object: %v\n", m)

	tID := m["tester_id"]
	fmt.Printf("\nstills for id(%s):\n", tID.(string))
	skills := m["tester_skills"]
	for _, skill := range skills.([]interface{}) {
		name := skill.(map[string]interface{})["skill_name"]
		level := skill.(map[string]interface{})["skill_level"]
		fmt.Printf("%s,%s\n", name.(string), level.(string))
	}
}

// template-01, parse string
func testGoTemplate01() {
	tmpl, err := template.New("test").Parse("hello, {{.}}\n")
	if err != nil {
		panic(err)
	}

	name := "vieira"
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		panic(err)
	}
}

// Inventory : struct used for template test
type Inventory struct {
	Material string
	Count    uint
}

// template-02, parse struct
func testGoTemplate02() {
	pattern := "{{.Count}} items are made of {{.Material}}\n"
	tmpl, err := template.New("test").Parse(pattern)
	if err != nil {
		panic(err)
	}

	sweaters := Inventory{"wool", 17}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

// template-03, multiple tmpls
func testGoTemplate03() {
	tagEn := "English"
	patternEn := "{{.Count}} items are made of {{.Material}}\n"
	tmpl, err := template.New(tagEn).Parse(patternEn)
	if err != nil {
		panic(err)
	}
	tagCn := "Chinese"
	patternCn := "{{ .Count }}个物料的材料是{{ .Material }}\n"
	tmpl, err = tmpl.New(tagCn).Parse(patternCn)
	if err != nil {
		panic(err)
	}

	sweaters := Inventory{"wool", 17}
	tmpl = tmpl.Lookup(tagEn)
	fmt.Println("\nCurrent template:", tmpl.Name())
	err = tmpl.ExecuteTemplate(os.Stdout, tagEn, sweaters)
	if err != nil {
		panic(err)
	}

	tmpl = tmpl.Lookup(tagCn)
	fmt.Println("\nCurrent template:", tmpl.Name())
	err = tmpl.ExecuteTemplate(os.Stdout, tagCn, sweaters)
	if err != nil {
		panic(err)
	}
}

// template-04, parse single file
func testGoTemplate04() {
	// filePath := os.Getenv("ZJGOPRJ") + "src/demo.hello/demos/data/tmpl_cn.txt"
	filePath := "src/demo.hello/demos/data/tmpl_cn.txt"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nCurrent template:", tmpl.Name())
	err = tmpl.Execute(os.Stdout, Inventory{"wool", 21})
	if err != nil {
		panic(err)
	}
}

// template-05, parse file with nested template
func testGoTemplate05() {
	fileSubTmpl := "src/demo.hello/demos/data/sub.tmpl"
	fileTmpl := "src/demo.hello/demos/data/tmpl_en.txt"
	tmpl, err := template.ParseFiles(fileSubTmpl, fileTmpl)
	if err != nil {
		panic(err)
	}

	// sub.tmpl + Inventory => tmpl_en.txt => stdout
	err = tmpl.ExecuteTemplate(os.Stdout, "main", Inventory{"wool", 21})
	if err != nil {
		panic(err)
	}
}

// template-06, parse files with nest template
func testGoTemplate06() {
	pattern := "src/demo.hello/demos/data/tmpl_*.txt"
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		panic(err)
	}
	fileSubTmpl := "src/demo.hello/demos/data/sub.tmpl"
	tmpl, err = tmpl.New("sub").ParseFiles(fileSubTmpl)
	if err != nil {
		panic(err)
	}

	name := tmpl.Lookup("main").Name()
	fmt.Printf("\n#1 template en(%s) output:\n", name)
	err = tmpl.ExecuteTemplate(os.Stdout, name, Inventory{"wool", 21})
	if err != nil {
		panic(err)
	}

	name = tmpl.Lookup("tmpl_cn.txt").Name()
	fmt.Printf("#2 template cn(%s) output:\n", name)
	err = tmpl.ExecuteTemplate(os.Stdout, name, Inventory{"wool", 27})
	if err != nil {
		panic(err)
	}
}

// read ini configs and create template
const confFile = "src/demo.hello/demos/data/test.conf"

func testCreateTemplateByConf() {
	fmt.Println("read configs and set template string")
	goInfos := readKVsFromConfigs(confFile, "default")
	fmt.Println("golang infos:")
	for k, v := range goInfos {
		fmt.Printf("%s=%s\n", k, v)
	}

	pattern := "\ngolang version {{.version}}, and bin path {{.path}}\n"
	tmpl, err := template.New("goInfos").Parse(pattern)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goInfos) // use map instead of struct
	if err != nil {
		panic(err)
	}
}

func readKVsFromConfigs(path, section string) map[string]string {
	// dependency: $ go get github.com/larspensjo/config
	cfg, err := config.ReadDefault(path)
	if err != nil {
		panic(fmt.Errorf("Fail to find %s, error: %s", path, err))
	}

	retMap := make(map[string]string)
	if cfg.HasSection(section) {
		if options, err := cfg.SectionOptions(section); err == nil {
			for _, option := range options {
				if value, err := cfg.String(section, option); err == nil {
					retMap[option] = value
				}
			}
		}
	}
	if len(retMap) == 0 {
		panic(fmt.Sprintf("no options found in section [%s]!\n", section))
	}
	return retMap
}

// read ini configs and create template file
func testCreateTemplateFile() {
	fmt.Println("\nsub.tmpl + test.conf => test_tmpl.conf => output.txt")

	// for sub tmpl, it supports diff data types, like array
	fileSubTmpl := "src/demo.hello/demos/data/sub.tmpl"
	fileTmpl := "src/demo.hello/demos/data/test_tmpl.conf"
	tmpl, err := template.ParseFiles(fileTmpl, fileSubTmpl)
	if err != nil {
		panic(err)
	}

	// for conf, it supports only key and value
	testInfos := readKVsFromConfigs(confFile, "test")

	const outPath = "src/demo.hello/demos/data/output.txt"
	fOutput, err := os.OpenFile(outPath, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fOutput.Close()

	err = tmpl.ExecuteTemplate(fOutput, "testInfos", testInfos)
	if err != nil {
		panic(err)
	}
	fmt.Println("create template:", outPath)
}

// MainUtils : main for utils
func MainUtils() {
	// testGetBaseFileName()
	// testCopyBytesToNull()

	// testURLParser()
	// testMD5Encode()
	// testHashFNV32()
	// testRegExp()

	// testMockGetRespBytesByRange()
	// testMockFileDownload()

	// testJSONMarshal()
	// testJSONUnmarshal01()
	// testJSONUnmarshal02()
	// testJSONUnmarshalToSlice()
	// testJSONUnmarshalToRawObj()

	// testGoTemplate01()
	// testGoTemplate02()
	// testGoTemplate03()
	// testGoTemplate04()
	// testGoTemplate05()
	// testGoTemplate06()

	// testCreateTemplateByConf()
	// testCreateTemplateFile()

	fmt.Println("golang utils demo DONE.")
}
