package main

import (
	"fmt"

	"demo.app/redis"
)

func main() {
	// redis env: docker run --name redis -p 6379:6379 --rm -d redis:4.0
	// ref: https://github.com/go-redis/redis/blob/master/example_test.go
	runN := 1

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
	fmt.Println("redis db demo done.")
}
