package engine

import (
	"gamersky/fetcher"
	"log"
)

type SimpleEngine struct {}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult := worker(r)

		requests = append(requests, parserResult.Requests...)

		for _,item := range parserResult.Items {
			log.Printf("Got item: %s: ", item)
		}
	}
}

func worker(r Request) ParserResult {
	body, err := fetcher.Get(r.Url)
	if err != nil {
		log.Printf("error fetching url %s: %v", r.Url, err)
	}

	return r.ParserFunc(body)
}