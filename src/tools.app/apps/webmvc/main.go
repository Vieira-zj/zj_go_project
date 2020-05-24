package main

import (
	"flag"
	"log"

	"tools.app/apps/webmvc/util"
)

func main() {
	// 解析配置文件
	// fpath := filepath.Join(os.Getenv("GOPATH"), "src/tools.app/apps/webmvc/app.dev.conf")
	fpath := flag.String("c", "app.dev.conf", "config file path")
	flag.Parse()

	config := util.NewConfigs()
	if _, err := config.Parse(*fpath); err != nil {
		log.Fatal("cofnigs parse error:", err.Error())
		return
	}
	config.WatchConfig()

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, os.Kill)
	// s := <-c
	// log.Println("Exit with signal:", s)
}
