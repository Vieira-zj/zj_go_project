package apps

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// const
const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 5 * time.Second
)

// Item :
type Item struct {
	Object     interface{}
	Expiration int64
}

// Expired : return true if the item expirated.
func (item Item) Expired() bool {
	if item.Expiration == -1 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

// Cache : use for outer struct.
type Cache struct {
	*cache
}

type cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
	mu                sync.RWMutex
	onEvicted         func(string, interface{}) // callback
	janitor           *janitor
}

/////////cache function

// Set : Add an item to the cache, replacing any existing item.
func (c *cache) Set(key string, object interface{}, d time.Duration) {
	if d == 0 {
		d = c.defaultExpiration
	}

	var expired int64
	if d > 0 {
		expired = time.Now().Add(d).UnixNano()
	}

	c.mu.Lock()
	c.items[key] = Item{
		Object:     object,
		Expiration: expired,
	}
	c.mu.Unlock()
}

func (c *cache) set(key string, object interface{}, d time.Duration) {
	if d == 0 {
		d = c.defaultExpiration
	}

	var expired int64
	if d > 0 {
		expired = time.Now().Add(d).UnixNano()
	}

	c.items[key] = Item{
		Object:     object,
		Expiration: expired,
	}
}

// Get : get an item from the cache.
// Returns the item or nil, and a bool indicating whether the key was found.
func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[key]
	if !found || item.Expired() {
		c.mu.RUnlock()
		return nil, false
	}
	c.mu.RUnlock()
	return item.Object, true
}

func (c *cache) get(key string) (interface{}, bool) {
	item, found := c.items[key]
	if !found || item.Expired() {
		return nil, false
	}
	return item.Object, true
}

// Add : add an item to the cache only if an item doesn't already exist for the given key,
// or if the existing item has expired. Returns an error otherwise.
func (c *cache) Add(key string, object interface{}, d time.Duration) error {
	c.mu.Lock()
	if _, found := c.get(key); found {
		c.mu.Unlock()
		err := fmt.Errorf("Item: %s has already exit", key)
		return err
	}
	c.set(key, object, d)
	c.mu.Unlock()
	return nil
}

// Replace: set a new value for the cache key only if it already exists,
// and the existing item hasn't expired. Returns an error otherwise.
func (c *cache) Replace(key string, object interface{}, d time.Duration) error {
	c.mu.Lock()
	if _, found := c.get(key); !found {
		c.mu.Unlock()
		err := fmt.Errorf("Item: %s dosen't exit", key)
		return err
	}
	c.set(key, object, d)
	c.mu.Unlock()
	return nil
}

//Increment : increment an item of number (int, TODO other type).
// Returns an error if the item's value is not an integer,
// if it was not found, or if it is not possible to increment it by n.
// To retrieve the incremented value, use one of the specialized methods, e.g. IncrementInt64.
func (c *cache) Increment(k string, n int64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("Item not found or expired")
	}

	switch v.Object.(type) {
	case int:
		v.Object = v.Object.(int) + int(n)
	default:
		c.mu.Unlock()
		return fmt.Errorf("not support value type")
	}

	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// Remove : remove an item from the cache. Does nothing if the key is not in the cache.
func (c *cache) Remove(key string) {
	c.mu.Lock()
	val, evicted := c.remove(key)
	c.mu.Unlock()
	if evicted {
		c.onEvicted(key, val)
	}
}

func (c *cache) remove(key string) (interface{}, bool) {
	if val, found := c.items[key]; found {
		delete(c.items, key)
		if c.onEvicted != nil {
			return val.Object, true
		}
	}
	return nil, false
}

type kv struct {
	key   string
	value interface{}
}

// DeleteExpired : delete all expired items from the cache.
func (c *cache) DeleteExpired() {
	var evictedItems []kv

	c.mu.Lock()
	for k, v := range c.items {
		if v.Expired() {
			v, evicted := c.remove(k)
			if evicted {
				evictedItems = append(evictedItems, kv{k, v})
			}
		}
	}
	c.mu.Unlock()

	for _, evicted := range evictedItems {
		c.onEvicted(evicted.key, evicted.value)
	}
}

// Items : return the item in the cache.
func (c *cache) Items() map[string]Item {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.items
}

// ItemCount : return the number of items in the cache. Equivalent to len(c.Items()).
func (c *cache) ItemCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// OnEvicted : sets an (optional) function that is called with the key and value
// when an item is evicted from the cache. (Including when it is deleted manually,
// but not when it is overwritten.) Set to nil to disable.
func (c *cache) OnEvicted(f func(string, interface{})) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.onEvicted = f
}

func (c *cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = map[string]Item{}
	// c.items = make(map[string]Item)
}

//////////// janitor function
type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func (j *janitor) Run(c *cache) {
	j.stop = make(chan bool)
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			fmt.Println("delete expired items")
			c.DeleteExpired()
		case <-j.stop:
			fmt.Println("stopped")
			ticker.Stop()
			return
		default:
			fmt.Println("sleep...")
			time.Sleep(time.Second)
		}
	}
}

////////////
func stopJanitor(c *cache) {
	c.janitor.stop <- true
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
	}
	c.janitor = j
	go j.Run(c)
}

func newCache(de time.Duration, m map[string]Item) *cache {
	if de <= 0 {
		de = NoExpiration
	}

	c := &cache{
		defaultExpiration: de,
		items:             m,
	}
	return c
}

func newCacheWithJanitor(defaultExpiration time.Duration, cleanupInterval time.Duration, items map[string]Item) *Cache {
	c := newCache(defaultExpiration, items)
	if cleanupInterval > 0 {
		runJanitor(c, cleanupInterval)
		runtime.SetFinalizer(c, stopJanitor)
	}
	return &Cache{c}
}

// New : return a new cache with a given default expiration duration and cleanup interval.
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)
	return newCacheWithJanitor(defaultExpiration, cleanupInterval, items)
}