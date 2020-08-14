package chanreq

import (
	"log"
)

//Dispatcher dispatcher
type Dispatcher struct {
	WorkerPool      chan chan Job
	NumberOfWorkers int
	jobQueue        chan Job
}

//NewDispatcher new dispatcher
func NewDispatcher(maxWorkers int, jobQueue chan Job) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, NumberOfWorkers: maxWorkers, jobQueue: jobQueue}
}

//Run runner
func (d *Dispatcher) Run() {
	for i := 0; i < d.NumberOfWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	log.Print("Dispatcher: Going to dispatch job")
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
