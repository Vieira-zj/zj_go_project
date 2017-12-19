package examples

import (
	"bufio"
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

// read and write file
func readFileTest() {
	path := filepath.Join(getExamplesDirPath(), "io_input.txt")
	values, err := readValues(path)
	if err != nil {
		fmt.Println("read file error:", err)
	} else {
		fmt.Println("file output:", values)
	}
}

func writeFileTest() {
	path := filepath.Join(getExamplesDirPath(), "io_output.txt")
	err := writeValues([]int{1, 11, 123, 1234}, path)
	if err != nil {
		fmt.Println("write file error:", err)
	}
}

func getExamplesDirPath() string {
	return filepath.Join(os.Getenv("ZJGO"), "src", "demo.hello", "examples")
}

func readValues(inFile string) (values []int, err error) {
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

func writeValues(values []int, outFile string) error {
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

// get the number of words
func countLineWordsTest() {
	path1 := filepath.Join(getExamplesDirPath(), "io_input.txt")
	path2 := filepath.Join(getExamplesDirPath(), "io_output.txt")
	paths := [2]string{path1, path2}
	counts := make(map[string]int)

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		defer f.Close()
		countLineWords(f, counts) // pass as reference
	}

	for word, num := range counts {
		fmt.Printf("%s\t%d\n", word, num)
	}
}

func countLineWords(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func deferTest() {
	defer func() {
		fmt.Println("self run function1 in defer")
	}()
	defer func() {
		fmt.Println("self run function2 in defer")
	}()
	defer myTrace("deferTest")()
	time.Sleep(3 * time.Second)
}

func myTrace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%.2f)", msg, time.Since(start).Seconds())
	}
}

func readExternalArgsTest() {
	var s1, s2 string
	for i := 1; i < len(os.Args); i++ {
		s1 += " " + os.Args[i]
	}
	fmt.Println("print args:", s1[1:])

	for _, arg := range os.Args[1:] {
		s2 += " " + arg
	}
	fmt.Println("print args:", s2[1:])

	fmt.Println("print args:", strings.Join(os.Args[1:], " "))
}

func flagTest() {
	// run cmd: $ ./main.go -period 3s
	period := flag.Duration("period", 1*time.Second, "sleep period")
	fmt.Printf("flag type: %T\n", period)

	flag.Parse()
	fmt.Printf("sleep for %v...\n", *period)
	time.Sleep(*period)
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
	hello("fname", "lastname")

	// readFileTest()
	// writeFileTest()
	// countLineWordsTest()

	// deferTest()

	// readExternalArgsTest()
	// flagTest()

	// fmtPrintfTest()
	// PrintGoEnvValues() // from hello.go

	fmt.Println("io demo.")
}
