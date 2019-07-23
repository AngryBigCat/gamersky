package scheduler

import "github.com/AngryBigCat/gamersky/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	var activeRequest engine.Request
	var activeWorker chan engine.Request

	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request

		for {

			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}

			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}

		}
	}()
}
