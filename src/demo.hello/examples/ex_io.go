package examples

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func isFileExists01(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func isFileExist02(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func testFileExist() {
	path := filepath.Join(getExamplesPath(), "io_input.txt")
	if isFileExist02(path) {
		fmt.Printf("file exist: %s\n", filepath.Base(path))
	} else {
		fmt.Printf("file not found: %s\n", filepath.Base(path))
	}
}

func getExamplesPath() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "demo.hello", "examples")
}

// io, writer and reader interface
func testIOReader() {
	// reader from string with 4 bytes
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				// should handle any remainding bytes
				fmt.Println("4 bytes string:", string(p[:n]))
				break
			}
			panic(err.Error())
		}
		fmt.Println("4 bytes string:", string(p[:n]))
	}
}

func testIOWriter() {
	// string into writer
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	var writer bytes.Buffer
	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			panic(err.Error())
		}
		if n != len(p) {
			panic("failed to write data!")
		}
	}
	fmt.Println("\nwriter string:\n", writer.String())
}

// Writer.Write(), Reader.Read()
func testIOWriterReader() {
	var buf bytes.Buffer // for rw
	buf.Write([]byte("\nwriter test\n"))
	buf.WriteTo(os.Stdout)

	fmt.Fprint(&buf, "writer test, and add buffer text")
	p := make([]byte, 4)
	var b []byte
	for {
		n, err := buf.Read(p)
		if err != nil {
			if err == io.EOF {
				b = append(b, p[:n]...)
				break
			}
			panic(err.Error())
		}
		b = append(b, p[:n]...)
	}
	fmt.Println("output bytes:", string(b))
}

// encode file content
func testIOEncodeText() {
	// create input stream
	in, err := os.Open(filepath.Join(getExamplesPath(), "io_input.txt"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer in.Close()

	// create output stream
	out, err := os.Create(filepath.Join(getExamplesPath(), "io_enc_output.txt"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer out.Close()
	base64Out := base64.NewEncoder(base64.StdEncoding, out)
	defer base64Out.Close()

	// data stream copy
	n, err := io.Copy(base64Out, in)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\nwrite %d bytes text to: %s\n", n, out.Name())
}

// read and write line from file
func testIOReadLine() {
	path := filepath.Join(getExamplesPath(), "io_input.txt")
	values, err := readIntValues(path)
	if err != nil {
		log.Fatalln("read file error:", err)
	} else {
		fmt.Println("\nfile text:", values)
	}
}

func readIntValues(filePath string) (values []int, err error) {
	// read file which each line is a number
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open input file %s\n", filePath)
		return
	}
	fmt.Printf("file type: %T\n", f)
	defer f.Close()

	br := bufio.NewReader(f)
	values = make([]int, 0)
	for {
		line, isPrefix, errRead := br.ReadLine()
		if errRead != nil {
			if errRead != io.EOF {
				err = errRead
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		value, errAtoi := strconv.Atoi(str)
		if errAtoi != nil {
			err = errAtoi
			return
		}
		values = append(values, value)
	}
	return
}

func testIOWriteLine() {
	path := filepath.Join(getExamplesPath(), "io_output.txt")
	err := writeIntValues([]int{1, 12, 123, 1234}, path)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("\nwrite to:", path)
}

func writeIntValues(values []int, outPath string) error {
	f, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", outPath)
		return err
	}
	defer f.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		f.WriteString(str + "\n")
	}
	return nil
}

// get total number of words in files
func testIOWordCount() {
	input1 := filepath.Join(getExamplesPath(), "io_input.txt")
	input2 := filepath.Join(getExamplesPath(), "io_output.txt")
	paths := [2]string{input1, input2}
	counts := make(map[string]int)

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		defer f.Close()
		countWords(f, counts) // pass as reference
	}

	fmt.Println("\nwords count:")
	for word, num := range counts {
		fmt.Printf("%s\t%d\n", word, num)
	}
}

func countWords(f *os.File, counts map[string]int) {
	// one word on each line
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

type myPoint struct {
	x int
	y int
}

func testFmtPrintf() {
	p := myPoint{1, 2}
	fmt.Println("\nprint format struct:")
	fmt.Printf("%v\n", p)  // {1 2}
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	fmt.Printf("%#v\n", p) // examples.myPoint{x:1, y:2}
	fmt.Printf("%T\n", p)  // examples.myPoint

	m := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("\nprint format map:")
	fmt.Printf("%v\n", m)
	fmt.Printf("%+v\n", m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("%T\n", m)
}

func testDeferFunc() {
	defer func() {
		fmt.Println("defer self-run func#1") // #3
	}()
	defer func() {
		fmt.Println("defer self-run func#2") // #2
	}()
	defer myTrace("deferTest")() // #1
	time.Sleep(3 * time.Second)
}

func myTrace(tag string) func() {
	start := time.Now()
	log.Printf("enter func#3: %s\n", tag)
	return func() {
		log.Printf("exit func#3: %s (%.2f)", tag, time.Since(start).Seconds())
	}
}

func testReadInputArgs() {
	// run cmd: ./gorun.sh main
	if len(os.Args) <= 1 {
		panic("no arguments pass in!")
	}

	var s1 string
	for i := 1; i < len(os.Args); i++ {
		s1 += " " + os.Args[i]
	}
	fmt.Println("\n#1 input arguments:", s1[1:])

	s2 := ""
	for _, arg := range os.Args[1:] {
		s2 += " " + arg
	}
	fmt.Println("#2 input args:", s2[1:])

	fmt.Println("#3 input args:", strings.Join(os.Args[1:], " "))
}

func testReadFlag01() {
	// run cmd: go run src/demo.hello/main/main.go -t 3
	waitTime := flag.Int("t", 1, "wait time by seconds")
	fmt.Printf("flag type: %T\n", waitTime)

	flag.Parse()
	fmt.Printf("sleep for %d ...\n", *waitTime)
	time.Sleep(time.Duration(*waitTime) * time.Second)
}

func testReadFlag02() {
	// run cmd: go run src/demo.hello/main/main.go -period 3s
	var duration time.Duration
	flag.DurationVar(&duration, "period", time.Second, "sleep period by seconds")
	fmt.Printf("\nflag type: %T\n", duration)

	flag.Parse()
	fmt.Printf("sleep for %v ...\n", duration)
	time.Sleep(duration)
}

func init() {
	fmt.Println("\n[ex_io.go] init.")
	// from ex_hello.go
	hello("fname", "lastname")
	PrintGoEnvValues()
}

// MainIO : main function for golang IO examples.
func MainIO() {
	// testFileExist()

	// testIOReader()
	// testIOWriter()
	// testIOWriterReader()

	// testIOEncodeText()
	// testIOReadLine()
	// testIOWriteLine()
	// testIOWordCount()

	// testFmtPrintf()
	// testDeferFunc()

	// testReadInputArgs()
	// testReadFlag01()
	// testReadFlag02()

	fmt.Println("golang io demo DONE.")
}
