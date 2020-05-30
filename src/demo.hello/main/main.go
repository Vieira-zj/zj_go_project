package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
)

func init() {
	fmt.Println("[main.go] init")
	fmt.Println("go version:", runtime.Version())
	fmt.Println("system arch:", runtime.GOARCH)
	fmt.Println("default int size:", strconv.IntSize)

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("\nruntime allocated memory: %d Kb\n", mem.Alloc/1024)
}

// flag test
var (
	code int
	msg  string
	help bool
)

func testFlagParser() {
	flag.IntVar(&code, "c", 200, "state code")
	flag.StringVar(&msg, "m", "hello world", "context message")
	flag.BoolVar(&help, "h", false, "help")

	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	fmt.Printf("arguments: code:%d, message:%s\n", code, msg)
}

func exitWithCtrlC() {
	// process blocked until ctrl-c signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Exit with signal:", s)
}

func main() {
	testFlagParser()
	// exitWithCtrlC()

	fmt.Println("GO main demo done.")
}
