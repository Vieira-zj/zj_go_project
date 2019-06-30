package redis

import (
	"errors"
	"fmt"
	"strings"
	"time"

	redis "github.com/go-redis/redis"
)

// TestRedis redis api general test.
type TestRedis struct {
	client *redis.Client
}

// NewTestRedis return a TestRedis instance.
func NewTestRedis() *TestRedis {
	options := redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&options)
	return &TestRedis{client: client}
}

// Close : close redis client.
func (t *TestRedis) Close() {
	t.client.Close()
}

// TestConnect test redis connection.
func (t TestRedis) TestConnect() {
	pong, err := t.client.Ping().Result()
	t.errHandler(err)
	fmt.Println("\nping:", pong)
}

// TestSetKV set a kv in redis.
func (t TestRedis) TestSetKV(key, val string) {
	err := t.client.Set(key, val, 0).Err()
	t.errHandler(err)
	fmt.Printf("set %s=%s\n", key, val)
}

// TestGetValue get a value in redis.
func (t TestRedis) TestGetValue(key string) {
	ret, err := t.client.Get(key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "nil") {
			fmt.Printf("get %s=nil\n", key)
		} else {
			panic(fmt.Sprintln("get value failed:", err))
		}
	} else {
		fmt.Printf("get %s=%s\n", key, ret)
	}
}

// TestSetKVByExpired set a value with expired time in redis.
func (t TestRedis) TestSetKVByExpired(key string, initN uint8) {
	err := t.client.Set(key, initN, time.Duration(2*time.Second)).Err()
	t.errHandler(err)

	if ret, err := t.client.Incr(key).Result(); err == nil {
		fmt.Printf("incr, and get %s=%d\n", key, ret)
	}
	if ret, err := t.client.IncrBy(key, 10).Result(); err == nil {
		fmt.Printf("incrby 10, and get %s=%d\n", key, ret)
	}

	time.Sleep(3 * time.Second)
	ret, err := t.client.Get(key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "nil") {
			fmt.Printf("key expired, get %s=nil\n", key)
		} else {
			panic(fmt.Sprintln("get value failed:", err))
		}
	} else {
		fmt.Printf("sleep, and get %s=%s\n", key, ret)
	}
}

// TestSetQueue set a value of list type in redis.
func (t TestRedis) TestSetQueue(key string) {
	for i := 0; i < 3; i++ {
		err := t.client.LPush(key, fmt.Sprintf("msg%d", i)).Err()
		t.errHandler(err)
	}

	len, err := t.client.RPush(key, "msgx").Result()
	t.errHandler(err)
	if len == 0 {
		panic(errors.New("queue is empty"))
	}
	fmt.Println("queue size:", len)

	if result, err := t.client.LRange(key, 0, len).Result(); err == nil {
		fmt.Println("queue items:", result)
	}
}

// TestDeleteKey delete a key in redis.
func (t TestRedis) TestDeleteKey(key string) {
	if err := t.client.Del(key).Err(); err != nil {
		t.errHandler(err)
	}
	fmt.Printf("success delete key=%s\n", key)
}

func (t TestRedis) errHandler(err error) {
	if err != nil {
		panic(err)
	}
}
