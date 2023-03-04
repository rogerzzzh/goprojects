package scheduler

import "goprojects/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {

}

func (s *SimpleScheduler) GenerateWorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}
