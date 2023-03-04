package scheduler

import (
	"goprojects/crawler/engine"
)

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (q *QueuedScheduler) GenerateWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	go func() {
		q.RequestChan <- request
	}()
}

func (q *QueuedScheduler) WorkerReady(c chan engine.Request) {
	go func() {
		q.WorkerChan <- c
	}()
}

func (q *QueuedScheduler) Run() {
	q.RequestChan = make(chan engine.Request)
	q.WorkerChan = make(chan chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-q.RequestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.WorkerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
