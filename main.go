package main

import (
	"gamersky/models"
	"gamersky/parser/gamersky"
	"gamersky/engine"
	"gamersky/scheduler"
	"fmt"
)

func main() {
	/*
	concurrentEngine.Run(engine.Request{
		Url:        "https://db2.gamersky.com/LabelJsonpAjax.aspx?callback=jQuery18303958816852369551_1535642187524&jsondata=%7B%22type%22%3A%22updatenodelabel%22%2C%22isCache%22%3Atrue%2C%22cacheTime%22%3A60%2C%22nodeId%22%3A%2211007%22%2C%22isNodeId%22%3A%22true%22%2C%22page%22%3A1%7D&_=1535642221589",
		ParserFunc: gamersky.ParserMain,
	})*/


	var newsList []models.News
	models.DB.Order("publish_at desc").Find(&newsList)

	var requests []engine.Request
	for _, v := range newsList {
		requests = append(requests, engine.Request{
			Url: v.Href,
			ParserFunc: func(bytes []byte) engine.ParserResult {
				return gamersky.ParserNews(bytes, v.Id)
			},
		})
	}

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	fmt.Println(len(requests))

	concurrentEngine.Run(requests...)
	/*for _,v := range newsList {
	}*/
}
