package main

import (
	"fmt"

	"demo.app/memcached"
	"demo.app/mongodb"
	"demo.app/redis"
)

func memcachedMain() {
	memcached.ConnectMemcacheAndTest()
}

func main() {
	isMongodbTest := false
	if isMongodbTest {
		// mongodb.ConnectToDbAndTest()

		mongodb.QueryBucketInfo()

		// mongodb.InsertRsRecords()
		// mongodb.InsertRsRecordsParallel()

		// cmd: ./main 10.200.30.11:8001
		// mongodb.PrintMongoOpLog()
	}

	isRedisTest := true
	if isRedisTest {
		redis.ConnectToRedisAndTest()
		// redis.MainRedis()
	}

	fmt.Println("data main done.")
}
