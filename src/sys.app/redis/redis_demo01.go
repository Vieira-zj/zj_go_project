package redis

import (
	"errors"
	"fmt"
	"strings"
	"time"

	redis "gopkg.in/redis.v5"
)

// ConnectToRedisAndTest : connect to redis, and test
// ref: https://github.com/go-redis/redis/blob/master/example_test.go
func ConnectToRedisAndTest() {
	options := redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&options)
	defer client.Close()

	handleErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	pong, err := client.Ping().Result()
	handleErr(err)
	fmt.Println(pong)

	// set val
	key := "1"
	val := "one"
	err = client.Set(key, val, 0).Err()
	handleErr(err)
	fmt.Printf("set %s=%s\n", key, val)

	// get val
	ret, err := client.Get(key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "nil") {
			fmt.Printf("get %s=nil\n", key)
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("get %s=%s\n", key, ret)
	}

	// set by expired, incr and incrby
	key = "test_total"
	err = client.Set(key, 8, time.Duration(2*time.Second)).Err()
	handleErr(err)

	if ret, err := client.Incr(key).Result(); err == nil {
		fmt.Printf("after incr 1, get %s=%d\n", key, ret)
	}
	if ret, err := client.IncrBy(key, 10).Result(); err == nil {
		fmt.Printf("after incrby 10, get %s=%d\n", key, ret)
	}

	time.Sleep(3 * time.Second)
	ret, err = client.Get(key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "nil") {
			fmt.Printf("expired, get %s=nil\n", key)
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("sleep, get %s=%s\n", key, ret)
	}

	// list
	key = "test_queue"
	for i := 0; i < 3; i++ {
		err = client.LPush(key, fmt.Sprintf("msg%d", i)).Err()
		handleErr(err)
	}
	len, err := client.RPush(key, "msgx").Result()
	handleErr(err)
	if len == 0 {
		panic(errors.New("queue is empty"))
	}

	fmt.Println("queue length:", len)
	if result, err := client.LRange(key, 0, len).Result(); err == nil {
		fmt.Println("queue items:", result)
	}
	if err := client.Del(key).Err(); err == nil {
		fmt.Println("success del:", key)
	}

	fmt.Println("redis test done.")
}
