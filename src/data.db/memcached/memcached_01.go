package memcached

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// ConnectMemcacheAndTest :
func ConnectMemcacheAndTest() {
	server := "127.0.0.1:11211"
	mc := memcache.New(server)
	if mc == nil {
		fmt.Println("memcache New failed")
	}
	// TODO:
	// https://blog.csdn.net/weixin_37696997/article/details/78760397
}
