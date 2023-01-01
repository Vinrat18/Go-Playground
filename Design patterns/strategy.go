package main

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     uint
	maxCapacity  uint
}

type Fifo struct{}
type Lru struct{}
type Lfu struct{}

func (*Fifo) evict(c *Cache) {
	fmt.Println("FIFO is being used for evict")
}

func (*Lru) evict(c *Cache) {
	fmt.Println("LRU is being used for evict")
}

func (*Lfu) evict(c *Cache) {
	fmt.Println("LFU is being used for evict")
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  100,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key string, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) string {
	return c.storage[key]
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {

	fifo := Fifo{}

	cache := initCache(&fifo)
	cache.add("1", "one")
	cache.add("2", "two")
	cache.evict()

	lru := Lru{}
	cache.setEvictionAlgo(&lru)
	cache.evict()
}
