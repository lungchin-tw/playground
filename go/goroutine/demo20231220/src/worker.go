package main

import (
	"fmt"
)

type Worker struct {
	id         string
	workerPool chan chan string
	jobChannel chan string
	quit       chan bool
}

func NewWorker(id string, workerPool chan chan string) *Worker {
	return &Worker{
		id:         id,
		workerPool: workerPool,
		jobChannel: make(chan string),
		quit:       make(chan bool),
	}
}

func (w *Worker) Run() {
	go func() {
		for {
			w.workerPool <- w.jobChannel
			select {
			case job := <-w.jobChannel:
				fmt.Printf("Worker(%v), Job(%v)\n", w.id, job)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
