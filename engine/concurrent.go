package engine

import (
	"fmt"
	"gamersky/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
} 

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 将in chan放到scheduler里， in在这块不进行操作
	in := make(chan Request)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	// 输出将在这里操作
	out := make(chan ParserResult)

	for i := 0; i < e.WorkerCount; i ++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		parserResult := <-out


		for _, r := range parserResult.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result

		}
	}()
}

func worker(r Request) (ParserResult, error){
	body, err := fetcher.Get(r.Url)
	fmt.Println(r.Url)
	if err != nil {
		log.Printf("error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}