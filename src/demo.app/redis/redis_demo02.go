package redis

import (
	"fmt"
	"sync"
	"time"

	redis "github.com/go-redis/redis"
)

// TestRedisOp redis general operations test cases.
type TestRedisOp struct {
	client *redis.Client
}

// NewTestRedisOp create a TestRedisOp instance.
func NewTestRedisOp() *TestRedisOp {
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
	fmt.Println("ping:", pong)

	return &TestRedisOp{client: client}
}

// Close : close redis client.
func (t *TestRedisOp) Close() {
	t.client.Close()
}

// StringOperation redis general operations for type string.
func (t TestRedisOp) StringOperation() {
	// #1
	const key1 = "name"
	err := t.client.Set(key1, "xys", 0).Err()
	defer t.DeleteKeyOp(key1)
	if err != nil {
		panic(err)
	}

	val, err := t.client.Get(key1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nget %s:%s\n", key1, val)

	// #2
	const key2 = "age"
	err = t.client.Set(key2, 20, time.Second).Err()
	defer t.DeleteKeyOp(key2)
	if err != nil {
		panic(err)
	}

	t.client.Incr(key2)
	t.client.Incr(key2)
	t.client.Decr(key2)
	val, err = t.client.Get(key2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nget %s=%s\n", key2, val)

	time.Sleep(time.Second)
	val, err = t.client.Get(key2).Result()
	if err != nil {
		fmt.Println("get failed:", err)
	} else {
		fmt.Printf("get %s=%s", key2, val)
	}
}

// ListOperation redis general operations for type list.
func (t TestRedisOp) ListOperation() {
	// #1
	const key = "fruit"
	t.client.RPush(key, "apple")
	t.client.LPush(key, "banana")
	defer t.DeleteKeyOp(key)
	len, err := t.client.LLen(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s size: %d\n", key, len)

	items, err := t.client.LRange(key, 0, len).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("all fruit: %v\n", items)

	// #2
	value, err := t.client.LPop(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nleft top fruit:", value)

	value, err = t.client.RPop(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("right top fruit:", value)
}

// SetOperation redis general set operations.
func (t TestRedisOp) SetOperation() {
	const (
		key1     = "blacklist"
		key2     = "whitelist"
		sameName = "the Elder"
	)

	// #1
	t.client.SAdd(key1, "Obama")
	t.client.SAdd(key1, "Hillary")
	t.client.SAdd(key1, sameName)
	defer t.DeleteKeyOp(key1)

	isMember, err := t.client.SIsMember(key1, "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nIs Bush in blacklist:", isMember)

	all, err := t.client.SMembers(key1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member of blacklist:", all)

	// #2
	t.client.SAdd(key2, sameName)
	defer t.DeleteKeyOp(key2)
	names, err := t.client.SInter(key1, key2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Inter results of blacklist and whitelist:", names)
}

// HashOperation redis hash general operations.
func (t TestRedisOp) HashOperation() {
	const key = "user_xys"
	t.client.HSet(key, "name", "xys")
	t.client.HSet(key, "age", "18")
	defer t.DeleteKeyOp(key)

	len, err := t.client.HLen(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nuser_xys fields count:", len)

	value, err := t.client.HGet(key, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user name:", value)
}

// ParallelConnsOp parallel do connections to redis.
func (t TestRedisOp) ParallelConnsOp() {
	const (
		routines = 10
		count    = 10
	)
	keys := make([]string, 0, routines*count)
	wg := sync.WaitGroup{}
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < count; j++ {
				key := fmt.Sprintf("name_%d", j)
				if err := t.client.Set(key, fmt.Sprintf("xys-%d-%d", i, j), 0).Err(); err != nil {
					fmt.Println("set failed:", err)
					continue
				}
				keys = append(keys, key)
			}
			fmt.Printf("Redis PoolStats: TotalConns=%d,IdleConns=%d\n",
				t.client.PoolStats().TotalConns, t.client.PoolStats().IdleConns)
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Second)

	defer func() {
		for _, key := range keys {
			t.DeleteKeyOp(key)
		}
	}()
}

// DeleteKeyOp delete a kv in redis.
func (t TestRedisOp) DeleteKeyOp(key string) {
	if len(key) == 0 {
		fmt.Println("delete key is empty!")
		return
	}

	if err := t.client.Del(key).Err(); err != nil {
		panic(err)
	}
	fmt.Printf("success delete key=%s\n", key)
}
