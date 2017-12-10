package main

import (
	"flag"
	"fmt"
	"os"

	"demo.tests/gotests"
)

// flags for Echo()
var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

func init() {
	fmt.Println("$GOROOT: " + os.Getenv("GOROOT"))
	fmt.Println("$GOPATH: " + os.Getenv("GOPATH"))
}

// cmd: go install src/demo.tests/main/test.go
func main() {
	flag.Parse()
	if err := gotests.Echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("test main done.")
}
