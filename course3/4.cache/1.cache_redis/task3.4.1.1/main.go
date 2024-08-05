package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Cacher interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type cache struct {
	client *redis.Client
}

func (c *cache) Set(key string, value interface{}) error {
	err := c.client.Set(key, value, 0).Err()
	return err
}

func (c *cache) Get(key string) (interface{}, error) {
	result, err := c.client.Get(key).Result()
	return result, err
}

func NewCache(client *redis.Client) Cacher {
	return &cache{
		client: client,
	}
}

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	cache := NewCache(client)

	err := cache.Set("some:key", "value")
	if err != nil {
		panic(err)
	}

	value, err := cache.Get("some:key")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
	user := &User{
		ID:   1,
		Name: "John",
		Age:  30,
	}

	err = cache.Set(fmt.Sprintf("user:%v", user.ID), user)
	if err != nil {
		panic(err)
	}

	value, err = cache.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
