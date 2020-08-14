package chanreq

import "log"

//Worker worker struct
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

//NewWorker new struct
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

//Start method starts the run loop for the worker, listening for a quit channel in case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Payload.RequestParser(); err != nil {
					log.Fatalf("Error crawling request: %s", err.Error())
				}

			case <-w.quit:
				return
			}
		}
	}()
}

//Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
