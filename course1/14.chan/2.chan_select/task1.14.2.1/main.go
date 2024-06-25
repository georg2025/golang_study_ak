package main

func main() {
}

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	case <-ch:
		return false
	default:
		return false
	}
}
