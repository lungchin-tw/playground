package main

import "strconv"

type Dispatcher struct {
	jobQueue   chan string
	workerPool chan chan string
	quit       chan bool
}

func NewDispatcher(jobQueue chan string, maxWorkersPool int) *Dispatcher {
	return &Dispatcher{
		jobQueue:   jobQueue,
		workerPool: make(chan chan string, maxWorkersPool),
		quit:       make(chan bool),
	}
}

func (d *Dispatcher) Run(maxWorkers int) {
	for i := 0; i < maxWorkers; i++ {
		NewWorker(strconv.Itoa(i), d.workerPool).Run()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func(job string) {
				worker := <-d.workerPool
				worker <- job
			}(job)

		case <-d.quit:
			return
		}
	}
}
