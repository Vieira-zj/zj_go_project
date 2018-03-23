package memcached

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// ConnectMemcacheAndTest : connect to memcached db, and test
func ConnectMemcacheAndTest() {
	// create a handle
	const server = "127.0.0.1:11211"
	mc := memcache.New(server)
	if mc == nil {
		fmt.Println("memcache New failed")
	}

	mc.Add(&memcache.Item{Key: "foo", Value: []byte("my value")})

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println("get failed, error:", err.Error())
	}
	if string(it.Key) == "foo" {
		fmt.Println("value is:", string(it.Value))
	}

	mc.Set(&memcache.Item{Key: "foo", Value: []byte("bluego")})
	it, err = mc.Get("foo")
	if err != nil {
		fmt.Println("add failed, error:", err.Error())
	}
	if string(it.Key) == "foo" {
		fmt.Println("value is:", string(it.Value))
	}

	// mc.Replace(&memcache.Item{Key: "foo", Value: []byte("mobike")})
	// it, err = mc.Get("foo")
	// if err != nil {
	// 	fmt.Println("replace failed, error:", err.Error())
	// }
	// if string(it.Key) == "foo" {
	// 	fmt.Println("value is:", string(it.Value))
	// }

	// delete an exist key
	err = mc.Delete("foo")
	if err != nil {
		fmt.Println("delete failed, error:", err.Error())
	} else {
		fmt.Println("delete success.")
	}

	// incrby
	err = mc.Set(&memcache.Item{Key: "myint", Value: []byte("1")})
	if err != nil {
		fmt.Println("set failed, error:", err.Error())
	}
	value, err := mc.Increment("myint", 7)
	if err != nil {
		fmt.Println("increment failed, error:", err.Error())
	}
	if string(it.Key) == "foo" {
		fmt.Println("after increment the value is:", value)
	}

	// decrby
	value, err = mc.Decrement("myint", 4)
	if err != nil {
		fmt.Println("decrement failed", err.Error())
	} else {
		fmt.Println("after decrement the value is:", value)
	}

	err = mc.Delete("myint")
	if err != nil {
		fmt.Println("delete failed, error:", err.Error())
	} else {
		fmt.Println("delete success.")
	}
}
