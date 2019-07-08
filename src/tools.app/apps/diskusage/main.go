package main

import (
	"flag"
	"fmt"

	mysvc "tools.app/services"
)

var (
	h = flag.Bool("h", false, "help")
	p = flag.String("p", "", "dir path")
)

// build cmd: ./gorun.sh tool diskusage
func main() {
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}

	diskUsage := mysvc.NewDiskUsage()
	if err := diskUsage.PrintDirDiskUsage(*p); err != nil {
		panic(err)
	}

	fmt.Println("tool disk usage done.")
}
