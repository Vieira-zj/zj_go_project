package demos

import (
	"fmt"
	"time"

	"demo.hello/cache"
)

// MainCache : main test for cache.go
func MainCache() {
	// Create a cache with a default expiration time of 5 minutes,
	// and which purges expired items every 3 seconds
	c := cache.New(5*time.Minute, 3*time.Second)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Get the string associated with the key "foo" from the cache
	if val, found := c.Get("foo"); found {
		fmt.Println(val)
	}

	// Set the value of the key "num" to 10, with the default expiration time.And add 1 to it.
	c.Set("num", 10, cache.DefaultExpiration)
	if err := c.Increment("num", 1); err != nil {
		fmt.Println(err)
	}
	if num, found := c.Get("num"); found {
		fmt.Println(num)
	}

	// Replace the value of item "foo"
	if err := c.Replace("foo", "change", cache.DefaultExpiration); err != nil {
		fmt.Println(err)
	}
	if val, found := c.Get("foo"); found {
		fmt.Println(val)
	}

	// Register callback function
	printDel := func(k string, v interface{}) {
		fmt.Printf("key=%s, value=%v\n", k, v)
	}
	c.OnEvicted(printDel)

	// Delete the item in the cache
	c.Delete("foo")
	if _, found := c.Get("foo"); !found {
		fmt.Println("deleted")
	}
}
