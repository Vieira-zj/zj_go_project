package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"tools.app/apps/webmvc/control"
	"tools.app/apps/webmvc/util"
)

func main() {
	// 解析配置文件
	dpath := filepath.Join(os.Getenv("GOPATH"), "src/tools.app/apps/webmvc/app.dev.conf")
	fpath := flag.String("c", dpath, "config file path")
	flag.Parse()

	config := util.NewConfigs()
	if _, err := config.Parse(*fpath); err != nil {
		log.Fatal("cofnigs parse error:", err.Error())
		return
	}
	config.WatchConfig()

	// 配置日志
	// 初始化数据库
	// 静态资源文件

	// 注册funcmap
	control.RegisterFuncMap()
	// 注册视图控制器
	control.RegisterView()
	// 注册控制器
	control.RegisterCtrl()

	serverConfig := config.GetSection("server")
	addr := serverConfig["server.addr"] + ":" + serverConfig["server.port"]
	log.Println("server listen at:", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(err.Error())
	}

	// 测试
	// curl -v "http://localhost:17891/index"
	// curl -v "http://localhost:17891/d/test"
	// curl -v -XPOST "http://localhost:17891/user/login" -H "Content-Type:application/json" \
	// -d '{"pagefrom":1,"pagesize":1,"asc":"field1","desc":"field2","id":111,"nickName":"test_user01","role":0,"code":"200"}'
}
