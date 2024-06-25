package main

import (
	"sync"
)

var mutex sync.RWMutex

func main() {
}

type Counter struct {
	count int64
}

func (c *Counter) Increment() {
	mutex.Lock()
	c.count++
	mutex.Unlock()
}

func (c *Counter) GetCount() int64 {
	mutex.RLock()
	result := c.count
	mutex.RUnlock()
	return result
}
