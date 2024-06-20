package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	number   int
	nextTask *task
}

func addTask(a int, q *queue) {
	newTask := task{number: a, nextTask: nil}
	if q.Tail == nil {
		q.Head = &newTask
		q.Tail = &newTask
	} else {
		q.Tail.nextTask = &newTask
		q.Tail = &newTask
	}
}

func takeTask(q *queue) (task, bool) {
	if q.Head == q.Tail {
		return *q.Head, false
	}
	currentTask := *q.Head
	q.Head = q.Head.nextTask
	return currentTask, true
}

type queue struct {
	Head *task
	Tail *task
}

func main() {
	newQueue := queue{Head: nil, Tail: nil}
	lastNumberNeeded := 100

	for i := 1; i <= lastNumberNeeded; i++ {
		addTask(i, &newQueue)
	}

	var wg sync.WaitGroup
	ch := make(chan int)

	go func() {
		defer close(ch)

		for {
			if newQueue.Head == nil {
				break
			}

			currentTask, stillHaveTasks := takeTask(&newQueue)
			wg.Add(1)
			go work(&currentTask, &wg, ch)

			if !stillHaveTasks {
				break
			}
		}
		wg.Wait()
	}()

	var answer int64 = 0
	for i := range ch {
		answer += int64(i)
	}

	fmt.Println("Summ of all squares from 1 to ", lastNumberNeeded, " is ", answer)
}

func work(task *task, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(1200 * time.Millisecond)
	result := task.number * task.number
	ch <- result
	wg.Done()
}
