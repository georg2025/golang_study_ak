package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mutex sync.RWMutex

type User struct {
	ID   int
	Name string
}

type Cache struct {
	cache map[string]*User
}

func NewCache() *Cache {
	data := make(map[string]*User)
	return &Cache{cache: data}
}

func (c *Cache) Set(key string, user *User) {
	mutex.Lock()
	c.cache[key] = user
	mutex.Unlock()
}

func (c *Cache) Get(key string) *User {
	mutex.RLock()
	result := c.cache[key]
	mutex.RUnlock()
	return result
}

func keyBuilder(keys ...string) string {
	var sb strings.Builder

	for _, i := range keys {
		sb.WriteString(i)
		sb.WriteString(":")
	}

	return strings.Trim(sb.String(), ":")
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		j := i

		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   j,
			Name: fmt.Sprint("user-", j),
		})
	}

	time.Sleep(5 * time.Millisecond)

	for i := 0; i < 100; i++ {
		j := i

		go func(j int) {
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(j))))
		}(j)
	}

	time.Sleep(5 * time.Millisecond)
}
