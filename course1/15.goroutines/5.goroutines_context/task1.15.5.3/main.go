package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	res := CallService()
	fmt.Println(res)
}

func CallService() string {
	data := make(chan string, 2)
	serviceLocator := NewServiceLocator()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer cancel()
		result, err := serviceLocator.SlowService(ctx)

		if err != nil {
			fmt.Println("error with SlowService: ", err)
		}

		data <- result
		fmt.Println("slow service done")
	}()

	go func() {
		defer cancel()
		result, err := serviceLocator.FastService(ctx)

		if err != nil {
			fmt.Println("error with SlowService: ", err)
		}

		data <- result
		fmt.Println("fast service done")
	}()

	<-ctx.Done()

	if len(data) > 1 {
		panic("error: more than one result")
	}

	checkService(serviceLocator)

	return <-data
}

func checkService(s *ServiceLocator) {
	if s.slow {
		panic("error: slow service called")
	}

	if !s.fast {
		panic("error: fast service not called")
	}
}

type ServiceLocator struct {
	client *http.Client
	fast   bool
	slow   bool
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (s *ServiceLocator) FastService(ctx context.Context) (string, error) {
	defer func() { s.fast = true }()
	time.Sleep(500 * time.Millisecond)
	return s.doRequest(ctx, "https://api.exmo.com/v1/ticker")
}

func (s *ServiceLocator) SlowService(ctx context.Context) (string, error) {
	defer func() { s.slow = true }()
	time.Sleep(2 * time.Second)
	return s.doRequest(ctx, "https://api.exmo.com/v1/ticker")
}

func (s *ServiceLocator) doRequest(ctx context.Context, url string) (string, error) {
	defer ctx.Done()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
