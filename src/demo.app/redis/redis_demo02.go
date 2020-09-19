package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// TestRedisOp redis general operations test cases.
type TestRedisOp struct {
	client *redis.Client
	ctx    context.Context
	cancel context.CancelFunc
}

// NewTestRedisOp create a TestRedisOp instance.
func NewTestRedisOp() *TestRedisOp {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})

	ctx, cancel := context.WithCancel(context.Background())

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ping:", pong)

	return &TestRedisOp{
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}
}

// Close : close redis client.
func (t *TestRedisOp) Close() {
	t.cancel()
	t.client.Close()
}

// StringOperation redis general operations for type string.
func (t *TestRedisOp) StringOperation() {
	// #1
	const key1 = "name"
	err := t.client.Set(t.ctx, "xys", 0, time.Duration(500)).Err()
	defer t.DeleteKeyOp(key1)
	if err != nil {
		panic(err)
	}

	val, err := t.client.Get(t.ctx, key1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nget %s:%s\n", key1, val)

	// #2
	const key2 = "age"
	err = t.client.Set(t.ctx, key2, 20, time.Second).Err()
	defer t.DeleteKeyOp(key2)
	if err != nil {
		panic(err)
	}

	t.client.Incr(t.ctx, key2)
	t.client.Incr(t.ctx, key2)
	t.client.Decr(t.ctx, key2)
	val, err = t.client.Get(t.ctx, key2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nget %s=%s\n", key2, val)

	time.Sleep(time.Second)
	val, err = t.client.Get(t.ctx, key2).Result()
	if err != nil {
		fmt.Println("get failed:", err)
	} else {
		fmt.Printf("get %s=%s", key2, val)
	}
}

// ListOperation redis general operations for type list.
func (t *TestRedisOp) ListOperation() {
	// #1
	const key = "fruit"
	t.client.RPush(t.ctx, key, "apple")
	t.client.LPush(t.ctx, key, "banana")
	defer t.DeleteKeyOp(key)
	len, err := t.client.LLen(t.ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s size: %d\n", key, len)

	items, err := t.client.LRange(t.ctx, key, 0, len).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("all fruit: %v\n", items)

	// #2
	value, err := t.client.LPop(t.ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nleft top fruit:", value)

	value, err = t.client.RPop(t.ctx, key).Result()
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
	t.client.SAdd(t.ctx, key1, "Obama")
	t.client.SAdd(t.ctx, key1, "Hillary")
	t.client.SAdd(t.ctx, key1, sameName)
	defer t.DeleteKeyOp(key1)

	isMember, err := t.client.SIsMember(t.ctx, key1, "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nIs Bush in blacklist:", isMember)

	all, err := t.client.SMembers(t.ctx, key1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member of blacklist:", all)

	// #2
	t.client.SAdd(t.ctx, key2, sameName)
	defer t.DeleteKeyOp(key2)
	names, err := t.client.SInter(t.ctx, key1, key2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Inter results of blacklist and whitelist:", names)
}

// HashOperation redis hash general operations.
func (t TestRedisOp) HashOperation() {
	const key = "user_xys"
	t.client.HSet(t.ctx, key, "name", "xys")
	t.client.HSet(t.ctx, key, "age", "18")
	defer t.DeleteKeyOp(key)

	len, err := t.client.HLen(t.ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nuser_xys fields count:", len)

	value, err := t.client.HGet(t.ctx, key, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user name:", value)
}

// ParallelConnsOp parallel do connections to redis.
func (t *TestRedisOp) ParallelConnsOp() {
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
				if err := t.client.Set(t.ctx, key, fmt.Sprintf("xys-%d-%d", i, j), 0).Err(); err != nil {
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
func (t *TestRedisOp) DeleteKeyOp(key string) {
	if len(key) == 0 {
		fmt.Println("delete key is empty!")
		return
	}

	if err := t.client.Del(t.ctx, key).Err(); err != nil {
		panic(err)
	}
	fmt.Printf("success delete key=%s\n", key)
}
