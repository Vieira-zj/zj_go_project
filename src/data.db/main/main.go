package main

import (
	"fmt"

	"data.db/redis"
)

// build cmd: $ GOOS=linux GOARCH=amd64 go build
// $ scp main qboxserver@10.200.20.21:~/zhengjin/main
func main() {

	// mongodb.ConnectToDbAndTest()
	// mongodb.InsertRecordsToRsDb()
	// mongodb.InsertRecordsToRsDbParallel()

	redis.ConnectToRedisAndTest()

	fmt.Println("db main done.")
}
