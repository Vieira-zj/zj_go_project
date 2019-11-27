// Disk tools: 1) print dir tree map; 2) print dir space usage.
//
// Build: ./gorun.sh tool diskusage
//
// Usage:
// Dir space usage: ./diskusage -u -p $(pwd)
// Print dir tree map: ./diskusage -t -p $(pwd) -l 2

package main

import (
	"flag"
	"fmt"

	mysvc "tools.app/services/diskusage"
	myutils "tools.app/utils"
)

var (
	h = flag.Bool("h", false, "help.")
	u = flag.Bool("u", false, "flag to print disk usage of dir, default path=cur_dir.")
	t = flag.Bool("t", false, "flag print dir tree map, default path=cur_dir and level=1.")
	l = flag.Int("l", 1, "level for dir tree map to print.")
	p = flag.String("p", myutils.GetCurPath(), "specified abs dir path.")
)

func main() {
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}

	fmt.Println("Run disk tools.")
	diskUsage := mysvc.NewDiskUsage()
	if *u {
		if err := diskUsage.PrintDirDiskUsage(*p); err != nil {
			panic(err)
		}
		return
	}
	if *t {
		if err := diskUsage.PrintFilesTree(*p, *l); err != nil {
			panic(err)
		}
	}
}
