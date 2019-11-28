package ddtest

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// DDArgs dd test arguments.
type DDArgs struct {
	Mode       int
	FileName   string
	BlockSize  int
	Count      int
	TimeoutMin int
}

// DDTest tests disk io perf by shell dd read and write.
type DDTest struct{}

// NewDDTest creates a DDTest instance.
func NewDDTest() *DDTest {
	return &DDTest{}
}

// DDCheck runs dd read and write to test disk io perf.
func (dd DDTest) DDCheck(args DDArgs) bool {
	base := "dd if=%s of=%s bs=%d count=%d oflag=direct"
	var cmd string
	if args.Mode == 1 { // rw
		base = fmt.Sprintf("%s iflag=direct", base)
		cmd = fmt.Sprintf(base, args.FileName, args.FileName+".out", args.BlockSize, args.Count)
	} else { // w
		cmd = fmt.Sprintf(base, "/dev/zero", args.FileName, args.BlockSize, args.Count)
	}

	// goroutine run dd
	ch := make(chan struct{})
	go func(ch chan<- struct{}) {
		defer func() {
			if p := recover(); p != nil {
				log.Println("run dd failed, error:", p.(error))
				close(ch)
			}
		}()

		results, err := dd.runShellCmd(cmd)
		if err != nil {
			panic(err)
		}
		log.Println(results)
		ch <- struct{}{}
	}(ch)
	time.Sleep(time.Second)

	// goroutine monitor dd
	chCheck := make(chan int64)
	go func(ch chan<- int64) {
		defer func() {
			if p := recover(); p != nil {
				log.Println("dd monitor failed, error:", p.(error))
				close(ch)
			}
		}()

		var lastSize int64
		for {
			curSize, err := dd.getFileSize(args.FileName)
			if err != nil {
				panic(err)
			}
			log.Println("current output file size:", curSize)
			if curSize == lastSize {
				ch <- curSize
				return
			}
			lastSize = curSize
			time.Sleep(time.Duration(3) * time.Second)
		}
	}(chCheck)

	select {
	case <-ch:
		log.Println("dd command done")
		fileSize, err := dd.getFileSize(args.FileName)
		if err != nil {
			panic(err)
		}
		if int64(args.BlockSize*args.Count) == fileSize {
			return true
		}
	case actualSize := <-chCheck:
		if int64(args.BlockSize*args.Count) == actualSize {
			return true
		}
	case <-time.After(time.Duration(args.TimeoutMin) * time.Minute):
		log.Println("dd run timeout!")
	}
	return false
}

func (dd DDTest) runShellCmd(shCommand string) (string, error) {
	log.Println("run shell command:", shCommand)
	var out bytes.Buffer
	sh := exec.Command("/bin/sh", "-c", shCommand)
	sh.Stdout = &out
	err := sh.Run()
	return out.String(), err
}

func (dd DDTest) getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
