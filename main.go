package main

// konsep caching

import (
	"fmt"
	// "math/rand"
	// "sync"
	"time"
)

type myService struct {
	cacher *cache
}

type cache struct {
	storage map[string]string
}

// layer repo
func (ms *myService) expensiveFunction(key string) string {
	time.Sleep(5 * time.Second)
	return fmt.Sprint("data-", key, ": Success")
}

// layer usecase
func (ms *myService) getData(key string) string {
	if ms.cacher != nil {
		if cacheData := ms.cacher.get(key); cacheData != "" {
			return cacheData
		}
	}

	result := ms.expensiveFunction(key)
	ms.cacher.set(key, result)

	return result
}

func (c *cache) set(key, value string) {
	c.storage[key] = value
}

func (c *cache) get(key string) string {
	v, ok := c.storage[key]
	if !ok {
		return ""
	}
	return v
}

func main() {
	cacher := &cache{
		storage: map[string]string{},
	}
	service := &myService{
		cacher: cacher,
	}

	key := "myData1"

	start := time.Now()
	fmt.Println("Calling expensive function")
	result := service.getData(key)
	fmt.Println("expensive function called, duration: ", time.Since(start))
	fmt.Println(result)

	start = time.Now()
	fmt.Println("Calling expensive function")
	result = service.getData(key)
	fmt.Println("cache expensive function called, duration: ", time.Since(start))
	fmt.Println(result)
}
