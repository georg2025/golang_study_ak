package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// hello. I am mutex. I am here to protect values from race condition
var mutex sync.RWMutex

// Hi. I am User struct
type User struct {
	ID   int
	Name string
}

// Hi. I am cashe. I am map, cause it is much easier to find, harder to lose and impossible to forget.
type Cache struct {
	cache map[string]interface{}
}

// Eeeehm... I am func for making a new cache
func NewCache() *Cache {
	data := make(map[string]interface{})
	return &Cache{cache: data}
}

// Me setting a value
func (c *Cache) Set(key string, user *User) {
	mutex.Lock()
	c.cache[key] = user
	mutex.Unlock()
}

// Me can get you a value. Actually I need this value as key to give you the same value as value
func (c *Cache) Get(key string) interface{} {
	mutex.RLock()
	result := c.cache[key]
	mutex.RUnlock()
	return result
}

// Me build keys from value
func keyBuilder(keys ...string) string {
	var sb strings.Builder

	for _, i := range keys {
		sb.WriteString(i)
		sb.WriteString(":")
	}

	return strings.Trim(sb.String(), ":")
}

// Me get User struct info
func GetUser(i interface{}) *User {
	return i.(*User)
}

// Me is main function. I am boss here. Obey
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

		go func(i int) {
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}

	time.Sleep(5 * time.Millisecond)
}
