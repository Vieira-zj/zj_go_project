package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"src/demo.tests/gotests"
)

func init() {
	fmt.Println("$GOROOT:", os.Getenv("GOROOT"))
	fmt.Println("$GOPATH:", os.Getenv("GOPATH"))
}

// flags for Echo()
var (
	h = flag.Bool("h", false, "help")
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

func main() {
	flag.Parse()
	fmt.Printf("\nflag [h] type: %T\n", h)
	if *h {
		flag.Usage()
	}

	gotests.Out = new(bytes.Buffer)
	fmt.Println("\ncmd arguments:", flag.Args())
	if err := gotests.Echo(*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("type: %T, results: %s\n", gotests.Out, gotests.Out.(*bytes.Buffer).String())

	fmt.Println("golang test main DONE.")
}
