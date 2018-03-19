package redis

import (
	"fmt"

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
	const key = "label"
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s=%s\n", key, val)

	// set key
	err = client.Set("desc", "fromgo", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis test done.")
}
