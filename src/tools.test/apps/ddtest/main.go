package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"tools.test/services"
)

func init() {
	fmt.Println("dd check tool run.")
}

var (
	fileName  string
	blockSize int
	count     int
	timeout   int
	help      bool
)

// cmd: ./ddtest -c 8 -bs 4
// scp ddtest root@cs50:/mnt/zjnfstest/nfstest
func main() {

	flag.StringVar(&fileName, "f", "test.file", "dd write file name.")
	flag.IntVar(&blockSize, "bs", 1, "block size (M).")
	flag.IntVar(&count, "c", 1, "block count.")
	flag.IntVar(&timeout, "t", 1, "timeout (minutes) for dd read and write file.")
	flag.BoolVar(&help, "h", false, "help.")

	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	fmt.Printf("timeout %d minutes.\n", timeout)

	args := services.DdArgs{
		FileName:   fileName,
		BlockSize:  blockSize * 1024 * 1024,
		Count:      count,
		TimeoutMin: timeout,
	}
	ok := services.TestDdCheck(args)
	if ok {
		log.Println("dd test done.")
		return
	}
	fmt.Printf("dd test failed for %v", args)
	os.Exit(-1)
}
