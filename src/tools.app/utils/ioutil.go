package utils

import "os"

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

// ReadFileContent reads file and returns file content string.
func ReadFileContent(path string) (string, error) {
	// TODO:
	return "", nil
}

// ReadFileLines reads file and returns all file lines.
func ReadFileLines(path string) ([]string, error) {
	// TODO:
	return []string{}, nil
}
