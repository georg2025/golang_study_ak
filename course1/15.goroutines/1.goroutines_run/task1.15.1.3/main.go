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

	for i := 1; i <= 100; i++ {
		addTask(i, &newQueue)
	}

	var wg sync.WaitGroup

	for {
		currentTask, stillHaveTasks := takeTask(&newQueue)
		if !stillHaveTasks {
			break
		}
		wg.Add(1)
		go work(&currentTask, &wg)
	}

	wg.Wait()
	fmt.Println("all tasks done")
}

func work(task *task, wg *sync.WaitGroup) {
	time.Sleep(1200 * time.Millisecond)
	result := task.number * task.number
	fmt.Println("The square of number ", task.number, " is ", result)
	wg.Done()
}
