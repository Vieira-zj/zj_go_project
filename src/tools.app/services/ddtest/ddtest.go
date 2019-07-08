package services

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// DDArgs : dd test arguments
type DDArgs struct {
	Mode       int
	FileName   string
	BlockSize  int
	Count      int
	TimeoutMin int
}

// DDTest test by shell dd command.
type DDTest struct{}

// NewDDTest create a DDTest instance.
func NewDDTest() *DDTest {
	return &DDTest{}
}

// DDCheck : run dd and check files
func (t DDTest) DDCheck(args DDArgs) bool {
	base := "dd if=%s of=%s bs=%d count=%d oflag=direct" // w
	var cmd string

	if args.Mode == 1 { // rw
		base = fmt.Sprintf("%s iflag=direct", base)
		cmd = fmt.Sprintf(base, args.FileName, args.FileName+".out", args.BlockSize, args.Count)
	} else {
		cmd = fmt.Sprintf(base, "/dev/zero", args.FileName, args.BlockSize, args.Count)
	}

	// routine run dd command
	chDD := make(chan bool)
	go func(ch chan<- bool) {
		output, err := t.runShellCmd(cmd)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
		ch <- true
	}(chDD)
	time.Sleep(time.Second)

	// routine montior dd command
	chCheck := make(chan int64)
	go func(ch chan<- int64) {
		var lastSize int64
		for {
			curSize := t.getFileSize(args.FileName)
			log.Println("cur file size:", curSize)
			if curSize == lastSize {
				ch <- curSize
				return
			}
			lastSize = curSize
			time.Sleep(time.Duration(5) * time.Second)
		}
	}(chCheck)

	select {
	case <-chDD:
		log.Println("dd command done")
		if int64(args.BlockSize*args.Count) == t.getFileSize(args.FileName) {
			return true
		}
	case actualSize := <-chCheck:
		if int64(args.BlockSize*args.Count) == actualSize {
			return true
		}
	case <-time.After(time.Duration(args.TimeoutMin) * time.Minute):
		log.Println("dd test timeout!")
	}
	return false
}

func (t DDTest) runShellCmd(shellCmd string) (string, error) {
	log.Println("run shell command:", shellCmd)
	sh := exec.Command("/bin/sh", "-c", shellCmd)
	var out bytes.Buffer
	sh.Stdout = &out
	err := sh.Run()

	return out.String(), err
}

func (t DDTest) getFileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		panic(err)
	}
	return fileInfo.Size()
}
