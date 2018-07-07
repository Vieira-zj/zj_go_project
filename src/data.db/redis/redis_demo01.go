package redis

import (
	"fmt"
	"strings"

	redis "gopkg.in/redis.v3"
)

// ConnectToRedisAndTest : connect to redis, and test
func ConnectToRedisAndTest() {
	options := redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	}
	client := redis.NewClient(&options)
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)

	// get key
	const key = "2"
	val, err := client.Get(key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "nil") {
			fmt.Printf("%s=nil\n", key)
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("%s=%s\n", key, val)
	}

	// set key
	err = client.Set("2", "two", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis test done.")
}
