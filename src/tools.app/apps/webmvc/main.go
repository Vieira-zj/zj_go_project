package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"src/tools.app/apps/webmvc/control"
	"src/tools.app/apps/webmvc/util"
)

func main() {
	// 解析配置文件
	dpath := filepath.Join(os.Getenv("GOPATH"), "src/tools.app/apps/webmvc/app.dev.conf")
	fpath := flag.String("c", dpath, "config file path")
	flag.Parse()

	config := new(util.Configs)
	if err := config.Parse(*fpath); err != nil {
		log.Fatal("cofnigs parse error:", err.Error())
		return
	}
	config.WatchConfig()

	// 服务配置
	serverConfig := config.GetSection("server")
	// 配置日志
	// 初始化数据库
	// 静态资源文件

	// 注册视图控制器
	control.RegisterView()
	// 注册控制器
	control.RegisterCtrl()

	addr := serverConfig["server.ip"] + ":" + serverConfig["server.port"]
	log.Println("server listen at:", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err.Error())
	}
}
