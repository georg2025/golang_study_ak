package main

import (
	"context"
	"time"
)

func main() {
	var res string
	res = contextWithDeadLine(context.Background(), 1*time.Second, 2*time.Second)
	println(res)
	res = contextWithDeadLine(context.Background(), 2*time.Second, 1*time.Second)
	println(res)
}

func contextWithDeadLine(ctx context.Context, contextDeadLine time.Duration, timeAfter time.Duration) string {
	ctx, cancel := context.WithTimeout(ctx, contextDeadLine)
	defer cancel()
	timer := time.NewTimer(timeAfter)

	select {
	case <-ctx.Done():
		return "context deadline exceeded"
	case <-timer.C:
		return "time after exceeded"
	}
}
