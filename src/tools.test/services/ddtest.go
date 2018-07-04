package services

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// DdArgs : dd test arguments
type DdArgs struct {
	FileName   string
	BlockSize  int
	Count      int
	TimeoutMin int
}

// TestDdCheck : run dd and check files
func TestDdCheck(args DdArgs) bool {
	chDd := make(chan bool)
	go func(ch chan<- bool) {
		cmd := fmt.Sprintf("dd if=/dev/zero of=%s bs=%d count=%d", args.FileName, args.BlockSize, args.Count)
		output, err := RunShellCmd(cmd)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
		ch <- true
	}(chDd)
	time.Sleep(time.Second)

	chCheck := make(chan int64)
	go func(ch chan<- int64) {
		var lastSize int64
		for {
			curSize := GetFileSize(args.FileName) // pending
			log.Println("cur file size:", curSize)
			if curSize == lastSize {
				ch <- curSize
				return
			}
			lastSize = curSize
			time.Sleep(8 * time.Second)
		}
	}(chCheck)

	select {
	case <-chDd:
		log.Println("dd cmd done")
		if int64(args.BlockSize*args.Count) == GetFileSize(args.FileName) {
			return true
		}
	case actualSize := <-chCheck:
		if int64(args.BlockSize*args.Count) == actualSize {
			return true
		}
	case <-time.After(time.Duration(args.TimeoutMin) * time.Minute):
		log.Println("dd timeout")
	}
	return false
}

// RunShellCmd : run shell commands
func RunShellCmd(shellCmd string) (string, error) {
	log.Println("run shell command:", shellCmd)
	cmd := exec.Command("/bin/sh", "-c", shellCmd)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	return out.String(), err
}

// GetFileSize : get file size
func GetFileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}
