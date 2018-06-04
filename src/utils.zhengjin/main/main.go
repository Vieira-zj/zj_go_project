package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	zjutils "utils.zhengjin/utils"
)

func init() {
	fmt.Println("kodo tool run.")
}

var (
	filePath string
	isEtag   bool
	help     bool
)

func flagParser() {
	flag.StringVar(&filePath, "f", "null", "file path")
	flag.BoolVar(&isEtag, "e", false, "get file etag")
	flag.BoolVar(&help, "h", false, "help")

	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if isEtag {
		printFileEtag()
	}
}

func printFileEtag() {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	strEtag, err := zjutils.GetEtagForText(string(content))
	if err != nil {
		panic(err)
	}
	fmt.Printf("get etag for file (%s): %s\n", filePath, strEtag)
}

// build cmd: /main$ GOOS=linux GOARCH=amd64 go build
// $ scp main qboxserver@10.200.20.21:~/zhengjin/main
// $ ./main -e -f test.file
func main() {
	flagParser()
}
