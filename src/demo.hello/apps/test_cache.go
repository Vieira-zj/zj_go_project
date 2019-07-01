package apps

import (
	"fmt"
	"time"
)

// TestCache test cache.go
func TestCache() {
	// Create a cache with a default expiration time of 5 minutes,
	// and which purges expired items every 3 seconds.
	c := New(time.Duration(5)*time.Minute, time.Duration(3)*time.Second)

	// Set the value of the key "foo" to "bar", with the default expiration time.
	const (
		tKey    = "foo"
		tVal    = "bar"
		tNewVal = "newBar"
		tKey2   = "num"
	)
	c.Set(tKey, tVal, DefaultExpiration)

	// Get the string associated with the key "foo" from the cache.
	if val, found := c.Get(tKey); found {
		fmt.Println("get:", val)
	}

	// Set value of the key "num" to 10, with the default expiration time. And add 1 to it.
	c.Set(tKey2, 10, DefaultExpiration)
	if err := c.Increment(tKey2, 1); err != nil {
		fmt.Println(err)
	}
	if num, found := c.Get(tKey2); found {
		fmt.Println("get incr number:", num)
	}

	// Replace the value of item "foo"
	if err := c.Replace(tKey, tNewVal, DefaultExpiration); err != nil {
		fmt.Println(err)
	}
	if val, found := c.Get(tKey); found {
		fmt.Println("replace new value:", val)
	}

	// Register callback function
	printDel := func(k string, v interface{}) {
		fmt.Printf("delete: key=%s, value=%v\n", k, v)
	}
	c.OnEvicted(printDel)

	// Delete the item in cache
	c.Remove(tKey)
	if _, found := c.Get(tKey); !found {
		fmt.Println("deleted")
	}

	time.Sleep(time.Duration(4) * time.Second)
	fmt.Println("cache test DONE.")
}
