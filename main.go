package main

import (
	"gamersky/engine"
	"gamersky/parser/news"
	"gamersky/scheduler"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(engine.Request{
		Url:        "https://db2.gamersky.com/LabelJsonpAjax.aspx?callback=jQuery18303958816852369551_1535642187524&jsondata=%7B%22type%22%3A%22updatenodelabel%22%2C%22isCache%22%3Atrue%2C%22cacheTime%22%3A60%2C%22nodeId%22%3A%2211007%22%2C%22isNodeId%22%3A%22true%22%2C%22page%22%3A1%7D&_=1535642221589",
		ParserFunc: news.ParserMain,
	})
}
