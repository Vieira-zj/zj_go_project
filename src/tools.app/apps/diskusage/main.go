package main

import (
	"flag"
	"fmt"

	mysvc "tools.app/services/diskusage"
)

var (
	h = flag.Bool("h", false, "help")
	t = flag.Bool("t", false, "print files tree of dir")
	l = flag.Int("l", 1, "files tree level to print")
	u = flag.Bool("u", false, "print disk usage of dir")
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
	// cmd: ./diskusage -u -p $(pwd)
	if *u {
		if err := diskUsage.PrintDirDiskUsage(*p); err != nil {
			panic(err)
		}
	}
	// cmd: ./diskusage -t -p $(pwd) -l 2
	if *t {
		if err := diskUsage.PrintFilesTree(*p, *l); err != nil {
			panic(err)
		}
	}

	fmt.Println("tool disk usage done.")
}
