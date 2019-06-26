package redis

import (
	"fmt"
	"sync"
	"time"

	redis "gopkg.in/redis.v3"
)

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)

	return client
}

func stringOperation(client *redis.Client) {
	const key1 = "name"
	err := client.Set(key1, "xys", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(key1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", val)

	const key2 = "age"
	err = client.Set(key2, 20, time.Second).Err()
	if err != nil {
		panic(err)
	}

	client.Incr(key2)
	client.Incr(key2)
	client.Decr(key2)
	val, err = client.Get(key2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("age:", val)

	time.Sleep(time.Second)
	val, err = client.Get("age").Result()
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("age:", val)
	}
}

func listOperation(client *redis.Client) {
	const key = "fruit"
	client.RPush(key, "apple")
	client.LPush(key, "banana")
	length, err := client.LLen(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("length: %d\n", length)

	items, err := client.LRange(key, 0, length).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("all fruit: %v\n", items)

	value, err := client.LPop(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit:", value)

	value, err = client.RPop(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit:", value)
}

func setOperation(client *redis.Client) {
	const key1 = "blacklist"
	client.SAdd(key1, "Obama")
	client.SAdd(key1, "Hillary")
	client.SAdd(key1, "the Elder")

	isMember, err := client.SIsMember(key1, "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist:", isMember)

	all, err := client.SMembers(key1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member:", all)

	const key2 = "whitelist"
	client.SAdd(key2, "the Elder")
	names, err := client.SInter(key1, key2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Inter result:", names)
}

func hashOperation(client *redis.Client) {
	client.HSet("user_xys", "name", "xys")
	client.HSet("user_xys", "age", "18")

	length, err := client.HLen("user_xys").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field count in user_xys:", length)

	value, err := client.HGet("user_xys", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user name:", value)
}

func connectPool(client *redis.Client) {
	const count = 10
	wg := sync.WaitGroup{}

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}
			// fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().FreeConns)
		}()
	}
	wg.Wait()
}

// MainRedis : main for redis test
func MainRedis() {
	client := createClient()
	defer client.Close()

	// stringOperation(client)
	// listOperation(client)
	// setOperation(client)
	// hashOperation(client)
	// connectPool(client)
}
