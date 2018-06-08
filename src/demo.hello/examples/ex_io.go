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

func isFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func getExamplesPath() string {
	return filepath.Join(os.Getenv("ZJGO"), "src", "demo.hello", "examples")
}

func testFileExist() {
	path := filepath.Join(getExamplesPath(), "io_input.log")
	if isFileExist(path) {
		fmt.Printf("file exist: %s\n", filepath.Base(path))
	} else {
		fmt.Printf("file not exist: %s\n", filepath.Base(path))
	}
}

// io, writer and reader interface
// reader
func testIOReader() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				// should handle any remainding bytes
				fmt.Println(string(p[:n]))
				break
			}
			panic(err.Error())
		}
		fmt.Println(string(p[:n]))
	}
}

// write
func testIOWriter() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}
	var writer bytes.Buffer

	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			panic(err.Error())
		}
		if n != len(p) {
			panic("failed to write data")
		}
	}
	fmt.Println(writer.String())
}

// Writer.Write(), Reader.Read()
func testIOWriterReader() {
	var buf bytes.Buffer // for rw
	buf.Write([]byte("writer test.\n"))
	buf.WriteTo(os.Stdout)
	fmt.Fprint(&buf, "writer test, add buffer text.")

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
	fmt.Println(string(b))
}

// encode file content
func encodeFileTextTest() {
	// build input stream
	f, err := os.Open(filepath.Join(getExamplesPath(), "io_input.txt"))
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// build output stream
	fEnc, err := os.Create(filepath.Join(getExamplesPath(), "io_enc_output.txt"))
	if err != nil {
		log.Fatalln(err)
	}
	defer fEnc.Close()

	// copy
	w := base64.NewEncoder(base64.StdEncoding, fEnc)
	io.Copy(w, f)
	w.Close()
}

// read and write line from file
func readLineTest() {
	path := filepath.Join(getExamplesPath(), "io_input.txt")
	values, err := readIntValues(path)
	if err != nil {
		fmt.Println("read file error:", err)
	} else {
		fmt.Println("file output:", values)
	}
}

func writeLineTest() {
	path := filepath.Join(getExamplesPath(), "io_output.txt")
	err := writeIntValues([]int{1, 11, 123, 1234}, path)
	if err != nil {
		fmt.Println("write file error:", err)
	}
}

func readIntValues(inFile string) (values []int, err error) {
	file, err := os.Open(inFile)
	if err != nil {
		fmt.Printf("Failed to open input file %s\n", inFile)
		return
	}
	fmt.Printf("file type: %T\n", file)
	defer file.Close()

	br := bufio.NewReader(file)
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

func writeIntValues(values []int, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Failed to create output file %s\n", outFile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

// get total number of words in files
func countFilesWordsTest() {
	path1 := filepath.Join(getExamplesPath(), "io_input.txt")
	path2 := filepath.Join(getExamplesPath(), "io_output.txt")
	paths := [2]string{path1, path2}
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

// examples
func deferFuncTest() {
	defer func() {
		fmt.Println("self run func#1 in defer") // #3
	}()
	defer func() {
		fmt.Println("self run func#2 in defer") // #2
	}()
	defer myTrace("deferTest")() // #1
	time.Sleep(3 * time.Second)
}

func myTrace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%.2f)", msg, time.Since(start).Seconds())
	}
}

func readPassArgsTest() {
	if len(os.Args) <= 1 {
		fmt.Println("no arguments pass in.")
		return
	}

	var s1 string
	for i := 1; i < len(os.Args); i++ {
		s1 += " " + os.Args[i]
	}
	fmt.Println("print args:", s1[1:])

	s2 := ""
	for _, arg := range os.Args[1:] {
		s2 += " " + arg
	}
	fmt.Println("print args:", s2[1:])

	fmt.Println("print args:", strings.Join(os.Args[1:], " "))
}

func readFlagTest1() {
	// run cmd: $ go run ./main.go -w 3
	waitTime := flag.Int("w", 1, "wait time by seconds")
	fmt.Printf("flag type: %T\n", waitTime)

	flag.Parse()
	fmt.Printf("sleep for %v...\n", *waitTime)
	time.Sleep(time.Duration(*waitTime) * time.Second)
	fmt.Println("flag test done.")
}

func readFlagTest2() {
	// run cmd: $ go run ./main.go -period 3s
	var duration time.Duration
	flag.DurationVar(&duration, "period", time.Second, "sleep period by seconds")
	fmt.Printf("flag type: %T\n", duration)

	flag.Parse()
	fmt.Printf("sleep for %v...\n", duration)
	time.Sleep(duration)
	fmt.Println("flag test done.")
}

// printf object
type myPoint struct {
	x int
	y int
}

func fmtPrintfTest() {
	p := myPoint{1, 2}
	fmt.Println("print struct with format.")
	fmt.Printf("%v\n", p)  // {1 2}
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	fmt.Printf("%#v\n", p) // examples.myPoint{x:1, y:2}
	fmt.Printf("%T\n", p)  // examples.myPoint

	m := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("\nprint map with format.")
	fmt.Printf("%v\n", m)
	fmt.Printf("%+v\n", m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("%T\n", m)
}

// MainIO : main function for IO examples.
func MainIO() {
	// from hello.go
	// hello("fname", "lastname")
	// PrintGoEnvValues()

	// testFileExist()
	// testIOReader()
	// testIOWriter()
	// testIOWriterReader()

	// encodeFileTextTest()
	// readLineTest()
	// writeLineTest()
	// countFilesWordsTest()

	// deferFuncTest()

	// readPassArgsTest()
	// readFlagTest1()
	// readFlagTest2()

	// fmtPrintfTest()

	fmt.Println("io demo.")
}
