package main

import (
	"gamersky/models"
	"gamersky/engine"
	"gamersky/scheduler"
	"gamersky/parser/gamersky"
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
		// 这里写成函数的原因是因为，如果不加func外层的话，会导致v.ID永远取最后一位
		requests = append(requests, func(news models.News) engine.Request {
			return engine.Request{
				Url: news.Href,
				ParserFunc: func(bytes []byte) engine.ParserResult {
					return gamersky.ParserNews(bytes, news.Id)
				},
			}
		}(v))
	}

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	concurrentEngine.Run(requests...)
}

