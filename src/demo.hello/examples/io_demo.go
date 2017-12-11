package examples

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func readArgsExamples() {
	var s1, s2 string
	for i := 1; i < len(os.Args); i++ {
		s1 += " " + os.Args[i]
	}
	fmt.Println("ex1 arguments:", s1[1:])

	for _, arg := range os.Args[1:] {
		s2 += " " + arg
	}
	fmt.Println("ex2 arguments:", s2[1:])

	fmt.Println("ex3 arguments:", strings.Join(os.Args[1:], " "))
}

func readValues(inFile string) (values []int, err error) {
	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println("Failed to open the input file ", inFile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outFile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func readFileTest() {
	values, err := readValues("./io_input.txt")
	if err != nil {
		fmt.Println("read file error:", err)
	} else {
		fmt.Println("file output:", values)
	}
}

func writeFileTest() {
	err := writeValues([]int{1, 11, 123, 1234}, "./io_output.txt")
	if err != nil {
		fmt.Println("write file error:", err)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func countLineTest() {
	paths := [2]string{"./io_input.txt", "./io_output.txt"}
	counts := make(map[string]int) // reference

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countLineTest: %v\n", err)
			continue
		}
		countLines(f, counts) // pass reference
		defer f.Close()
	}

	for line, num := range counts {
		fmt.Printf("%s\t%d\n", line, num)
	}
}

func deferTest() {
	defer myTrace("deferTest")()
	time.Sleep(5 * time.Second)
	return
}

func myTrace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%.2f)", msg, time.Since(start).Seconds())
	}
}

func flagTest() {
	// run cmd: $ ./io_demos -period 5s
	period := flag.Duration("period", 1*time.Second, "sleep period")

	flag.Parse()
	fmt.Printf("sleep for %v...", *period)
	time.Sleep(*period)
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
	// readArgsExamples()

	// readFileTest()
	// writeFileTest()

	// countLineTest()

	// deferTest()
	// flagTest()

	fmtPrintfTest()

	// invoke function of "io.demo.go"
	// PrintGoEnvValues()

	fmt.Println("io demo.")
}
