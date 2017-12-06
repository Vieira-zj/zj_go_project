package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strconv"

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

func mainIO() {
	values, err := readValues("./io_input.txt")
	if err != nil {
		fmt.Println("read file error:", err)
	} else {
		fmt.Println("file output:", values)
	}

	err1 := writeValues([]int{1, 11, 123, 1234}, "./io_output.txt")
	if err1 != nil {
		fmt.Println("write file error:", err)
	}

	fmt.Println("io demo.")
}
