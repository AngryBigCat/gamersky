package engine

import "fmt"

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
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("%v\n", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <- in
			result := worker(request)
			out <- result
		}
	}()
}
