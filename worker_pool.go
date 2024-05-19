package workerpool

import (
	"sync"
)

type TaskRunner interface {
	RunTask()
}

type WorkerPool struct {
	tasks       []TaskRunner
	taskQueue   chan TaskRunner
	concurrency int
	wg          sync.WaitGroup
}

func NewWorkerPool(tasks []TaskRunner, concurrency int) *WorkerPool {
	return &WorkerPool{
		tasks:       tasks,
		concurrency: concurrency,
	}
}

func (wp *WorkerPool) worker() {
	// Dequeue  task
	for task := range wp.taskQueue {
		task.RunTask()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Start() {
	// Start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// Enqueue task
	wp.wg.Add(len(wp.tasks))
	wp.taskQueue = make(chan TaskRunner, len(wp.tasks))
	for _, task := range wp.tasks {
		wp.taskQueue <- task
	}
	close(wp.taskQueue)

	wp.wg.Wait()
}
