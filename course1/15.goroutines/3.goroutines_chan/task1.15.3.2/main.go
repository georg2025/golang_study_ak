package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "https://httpbin.org/get"
	parallelRequest := 5
	requestCount := 50
	result := benchRequest(url, parallelRequest, requestCount)

	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("error with request sending: %d", statusCode))
		}
	}

	fmt.Println("All goroutines done")
}

func benchRequest(url string, parallelRequest int, requestCount int) <-chan int {
	semaphore := make(chan struct{}, parallelRequest)
	ch := make(chan int)
	var wg sync.WaitGroup

	go func() {
		defer close(ch)
		for i := 0; i < requestCount; i++ {
			semaphore <- struct{}{}
			wg.Add(1)
			go func() {
				statusCode, _ := httpRequest(url)
				ch <- statusCode
				<-semaphore
				wg.Done()
			}()
		}

		wg.Wait()
	}()

	return ch
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, nil
}
