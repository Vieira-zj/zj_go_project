package cache

import (
	"time"
)

// const
const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

// Item :
type Item struct {
	Object     interface{}
	Expiration int64
}

// Expired : return true if the item expirated
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

// Cache : use for outer struct
type Cache struct {
	*cache
}

type cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
	onEvicted         func(string, interface{}) // callback
	// janitor           *janitor
}

// TODO:
