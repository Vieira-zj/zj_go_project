package memcached

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// ConnectMemcacheAndTest : connect to memcached db, and test op
func ConnectMemcacheAndTest() {
	// create connection
	const server = "127.0.0.1:11211"
	mc := memcache.New(server)
	if mc == nil {
		panic("memcache New failed")
	}

	// demo #1
	var (
		tKey      = "foo"
		tValue    = "my_value"
		tNewValue = "bluego"
	)

	// add kv
	if err := mc.Add(&memcache.Item{Key: tKey, Value: []byte(tValue)}); err != nil {
		panic(fmt.Sprintln("add kv failed:", err))
	}

	// get kv
	it, err := mc.Get(tKey)
	if err != nil {
		panic(fmt.Sprintln("get kv failed:", err))
	}
	if string(it.Key) == tKey {
		fmt.Printf("get: %s=%s", tKey, string(it.Value))
	}

	// update kv
	if err := mc.Set(&memcache.Item{Key: tKey, Value: []byte(tNewValue)}); err != nil {
		panic(fmt.Sprintln("set kv failed:", err))
	}
	it, err = mc.Get(tKey)
	if err != nil {
		panic(fmt.Sprintln("get kv failed:", err))
	}
	if string(it.Key) == tKey {
		fmt.Printf("update and get: %s=%s", tKey, string(it.Value))
	}

	// delete key
	if err := mc.Delete(tKey); err != nil {
		panic(fmt.Sprintln("delete kv failed:", err))
	}
	fmt.Printf("delete key=%s success\n", tKey)

	// demo#2
	var (
		incrKey        = "myint"
		initN          = "1"
		incrN   uint64 = 7
		decrN   uint64 = 6
	)

	// incrby
	if err := mc.Set(&memcache.Item{Key: incrKey, Value: []byte(initN)}); err != nil {
		panic(fmt.Sprintln("set kv failed:", err))
	}
	value, err := mc.Increment(incrKey, incrN)
	if err != nil {
		panic(fmt.Sprintln("increment failed:", err))
	}
	fmt.Println("increment value:", value)

	// decrby
	value, err = mc.Decrement(incrKey, decrN)
	if err != nil {
		panic(fmt.Sprintln("decrement failed:", err))
	}
	fmt.Println("decrement value:", value)

	// delete key
	err = mc.Delete(incrKey)
	if err != nil {
		panic(fmt.Sprintln("delete key failed:", err))
	}
	fmt.Printf("delete key=%s success\n", incrKey)
}
