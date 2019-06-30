package main

import (
	"flag"
	"fmt"
	"os"

	"tools.app/services"
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
	flag.IntVar(&mode, "m", 0, "test mode: 0-w, 1-rw.")
	flag.StringVar(&fileName, "f", "test.file", "dd write file name.")
	flag.IntVar(&blockSize, "bs", 1, "block size (M).")
	flag.IntVar(&count, "c", 1, "block count.")
	flag.IntVar(&timeout, "t", 1, "timeout (minutes) for dd read and write file.")
	flag.BoolVar(&help, "h", false, "help.")
}

// cmd: ./ddtest -c 8 -bs 4
// scp ddtest root@cs50:/mnt/zjnfstest/nfstest
func main() {
	fmt.Println("dd check tool started.")
	flagInit()
	flag.Parse()

	if help {
		flag.Usage()
		return
	}
	fmt.Printf("ddtest timeout %d(minutes)\n", timeout)

	ddtest := services.NewDDTest()
	args := services.DDArgs{
		Mode:       mode,
		FileName:   fileName,
		BlockSize:  blockSize * 1024 * 1024,
		Count:      count,
		TimeoutMin: timeout,
	}

	ok := ddtest.DDCheck(args)
	if ok {
		fmt.Println("dd test pass")
	} else {
		fmt.Printf("dd test failed for %v!", args)
		os.Exit(99)
	}
}
