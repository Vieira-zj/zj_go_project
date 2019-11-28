// Test disk io perf by dd read and write.
//
// Build: ./gorun.sh tool ddtest
//
// Usage:
// ./ddtest -c 8 -bs 4
// scp ddtest root@cs50:/mnt/zjnfstest/nfstest

package main

import (
	"flag"
	"fmt"
	"os"

	mysvc "tools.app/services/ddtest"
)

var (
	mode      int
	fileName  string
	blockSize int
	count     int
	timeout   int
	help      bool
)

func flagInit() {
	flag.BoolVar(&help, "h", false, "help.")
	flag.IntVar(&mode, "m", 0, "dd disk io test mode: 0-w and 1-rw. default mode=0.")
	flag.StringVar(&fileName, "f", "ddtest.out", "dd output file name.")
	flag.IntVar(&blockSize, "bs", 1, "dd output file block size (M).")
	flag.IntVar(&count, "c", 1, "dd output file blocks count.")
	flag.IntVar(&timeout, "t", 1, "dd disk io test timeout (minutes).")
}

func main() {
	fmt.Println("dd disk io test tool started.")
	flagInit()
	flag.Parse()

	if help {
		flag.Usage()
		return
	}
	fmt.Printf("dd disk io test timeout: %d(minutes)\n", timeout)

	args := mysvc.DDArgs{
		Mode:       mode,
		FileName:   fileName,
		BlockSize:  blockSize * 1024 * 1024,
		Count:      count,
		TimeoutMin: timeout,
	}

	ddtest := mysvc.NewDDTest()
	ok := ddtest.DDCheck(args)
	if ok {
		fmt.Println("dd disk io perf test Pass")
	} else {
		fmt.Println("dd disk io perf test Failed, test args:", args)
		os.Exit(99)
	}
}
