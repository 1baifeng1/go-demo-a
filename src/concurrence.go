package main

import (
	"fmt"
	"runtime"
	"time"
)

type Job interface {
	Do()
}

type Worker struct {
	JobQueue chan Job
}

func NewWorker() Worker {
	return Worker{JobQueue: make(chan Job)}
}

func (w Worker) Run(wq chan chan Job) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

type WorkerPool struct {
	WorkerNum   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerNum int) *WorkerPool {
	return &WorkerPool{
		WorkerNum:   workerNum,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerNum),
	}
}

func (wp *WorkerPool) Run() {
	fmt.Println("Initing worker.")
	for i := 0; i < wp.WorkerNum; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}

	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}

type Score struct {
	Num int
}

func (s *Score) Do() {
	fmt.Println("num:", s.Num)
	time.Sleep(1 * 1 * time.Second)
}

func main() {
	num := 2 * 10000
	p := NewWorkerPool(num)
	p.Run()

	dataNum := 100 * 10000
	go func() {
		for i := 0; i < dataNum; i++ {
			sc := &Score{i}
			p.JobQueue <- sc
		}
	}()

	for {
		fmt.Println("runtime.NumGoroutine():", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
