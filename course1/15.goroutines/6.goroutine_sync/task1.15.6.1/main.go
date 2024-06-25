package main

import (
	"fmt"
	"strings"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	var sb strings.Builder
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for _, i := range goroutines {
		wg.Add(1)
		j := i

		go func() {
			defer wg.Done()
			info := j()
			mutex.Lock()
			sb.WriteString(info)
			sb.WriteString("\n")
			mutex.Unlock()
		}()
	}

	wg.Wait()
	return strings.Trim(sb.String(), "\n")
}

func main() {
	count := 1000
	goroutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}

	fmt.Println(waitGroupExample(goroutines...))
}
