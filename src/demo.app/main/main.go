package main

import (
	"fmt"
	"os"

	"demo.app/memcached"
	"demo.app/mongodb"
	"demo.app/redis"
)

func memcachedMain() {
	memcached.ConnectMemcacheAndTest()
}

func mongodbMain(runN uint8) {
	// #1 mongo test
	if runN == 1 {
		mongodb.ConnectToDbAndTest()
	}
	// #2 bucket
	if runN == 2 {
		const (
			bucket = "test_bucket_transfer_data11"
			uid    = 1380469264
		)
		bucketInfo := mongodb.NewBucketInfo(bucket, uid)
		bucketInfo.QueryBucketInfo()
	}
	// #3 rs
	if runN == 3 {
		rsOp := mongodb.NewRsOperation()
		defer rsOp.Close()
		rsOp.InsertRsRecords()
		// rsOp.InsertRsRecordsParallel()
	}
	// #4 mongo op logs
	// cmd: ./main 10.200.30.11:8001
	if runN == 4 {
		mgoOp := mongodb.NewMgoOpertion(os.Args[1])
		defer mgoOp.Close()
		mgoOp.PrintMgoOpLogs()
	}
}

func redisMain(runN int) {
	// redis env: docker run --name redis -p 6379:6379 --rm -d redis:4.0
	// ref: https://github.com/go-redis/redis/blob/master/example_test.go

	demo01 := func() {
		r := redis.NewTestRedis()
		defer r.Close()
		r.TestConnect()

		const (
			key   = "1"
			val   = "one"
			cKey  = "test_count"
			initN = 3
			qKey  = "test_queue"
		)
		r.TestSetKV(key, val)
		r.TestGetValue(key)
		r.TestSetKVByExpired(cKey, initN)
		r.TestSetQueue(qKey)
		defer func() {
			r.TestDeleteKey(key)
			r.TestDeleteKey(cKey)
			r.TestDeleteKey(qKey)
		}()
	}

	demo02 := func() {
		op := redis.NewTestRedisOp()
		defer op.Close()

		op.StringOperation()
		op.ListOperation()

		op.SetOperation()
		op.HashOperation()
	}

	demo03 := func() {
		op := redis.NewTestRedisOp()
		defer op.Close()
		op.ParallelConnsOp()
	}

	if runN == 1 {
		demo01()
	}
	if runN == 2 {
		demo02()
	}
	if runN == 3 {
		demo03()
	}
}

func main() {
	// memcachedMain()
	// mongodbMain(1)
	redisMain(1)

	fmt.Println("memory db demo DONE.")
}
