package engine

import (
	"goprojects/crawler/deduplicator"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemSaver   chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	GenerateWorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	requestDuplicator := deduplicator.MapDeduplicator{}
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler, e.Scheduler.GenerateWorkerChan(), out)
	}

	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	for result := range out {
		log.Printf("Concurrent Engine Info: got items %s", result.Items)
		for _, item := range result.Items {
			//log.Printf("Concurrent Engine Info: got %s", item)
			go func(item interface{}) {
				e.ItemSaver <- item
			}(item)
		}

		for _, newRequest := range result.Requests {
			if requestDuplicator.IsExist(newRequest.Url) {
				log.Printf("Concurrent Engine Info: Duplicated URL %s\n", newRequest.Url)
				continue
			}
			requestDuplicator.Add(newRequest.Url)
			e.Scheduler.Submit(newRequest)
		}
	}
}

func createWorker(ready ReadyNotifier, in chan Request, out chan ParseResult) {
	go func() {
		ready.WorkerReady(in)
		for request := range in {
			result, err := worker(request)
			if err != nil {
				log.Printf("Worker Error: worker error %v", err)
				continue
			}
			out <- result
			ready.WorkerReady(in)
		}
	}()
}
