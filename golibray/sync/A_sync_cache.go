package syncTest

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Cache struct {
	data   map[string]string
	rmlock sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) string {
	c.rmlock.RLock()
	defer c.rmlock.Unlock()
	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.rmlock.Lock()
	defer c.rmlock.Unlock()
	c.data[key] = value
}

func ReadRequest(cache *Cache, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := cache.Get(key)
		log.Printf("Read request - Key:%s,Value:%s\n", key, value)
		time.Sleep(2 * time.Second)
	}
}

func WriteRequest(cache *Cache, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Set(key, value)
		fmt.Printf("Write request - Key: %s, Value: %s\n", key, value)
		time.Sleep(2 * time.Second)
	}
}
