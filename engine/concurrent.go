package engine

import (
	"fmt"
	"log"

	"github.com/AngryBigCat/gamersky/fetcher"
)

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	Run()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 输入channel
	in := make(chan Request)

	// e.Scheduler.Run()
	// 将输入channel配置进Scheduler
	// e.Scheduler.ConfigureMasterWorkerChan(in)
	// 输出channel
	out := make(chan ParserResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 1
	for {
		parserResult := <-out

		for _, item := range parserResult.Items {
			fmt.Printf("Go item #%d: %v \n", itemCount, item)
			itemCount++
		}

		for _, r := range parserResult.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Get(r.Url)
	if err != nil {
		log.Printf("error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
