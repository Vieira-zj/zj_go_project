package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// IsFileExist returns bool for file exist.
func IsFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DeleteFile delete a file or an empty directory.
func DeleteFile(path string) error {
	exist, err := IsFileExist(path)
	if !exist {
		return nil
	}
	if err != nil {
		return err
	}
	return os.Remove(path)
}

// GetCurPath returns current run abs path.
func GetCurPath() string {
	dir, _ := filepath.Split(os.Args[0])
	return dir
}

// ********* File Read

// ReadFileContent reads file and returns file content string.
func ReadFileContent(path string) (string, error) {
	exist, err := IsFileExist(path)
	if !exist {
		return "", fmt.Errorf("file (%s) not found", path)
	}
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadFileContentBuf reads file with buffer and returns file content bytes.
func ReadFileContentBuf(path string) ([]byte, error) {
	var (
		bufSize  = 1024
		retBytes = make([]byte, 0, bufSize*10)
	)

	exist, err := IsFileExist(path)
	if !exist {
		return retBytes, fmt.Errorf("file (%s) not found", path)
	}
	if err != nil {
		return retBytes, err
	}

	f, err := os.Open(path)
	if err != nil {
		return retBytes, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		buf := make([]byte, bufSize)
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		retBytes = append(retBytes, buf...)
	}
	return retBytes, nil
}

// ReadFileLines reads file and returns all file lines.
func ReadFileLines(path string) ([]string, error) {
	lines := make([]string, 0, 100)
	exist, err := IsFileExist(path)
	if !exist {
		return lines, fmt.Errorf("file (%s) not found", path)
	}
	if err != nil {
		return lines, err
	}

	f, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// line, _, err := r.ReadLine()
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return lines, err
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

// ********* File Write

// WriteContentToFile writes content to file.
func WriteContentToFile(path, content string, isOverwrite bool) error {
	exist, err := IsFileExist(path)
	if err != nil {
		return err
	}
	if !isOverwrite && exist {
		return fmt.Errorf("file (%s) is exist", path)
	}

	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

// AppendContentToFile if file exist, appends content at the end, or writes content to a new file.
func AppendContentToFile(path, content string) (int, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.WriteString(content)
}
