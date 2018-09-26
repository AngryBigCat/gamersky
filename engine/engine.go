package engine

import (
	"gamersky/fecher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fecher.Get(r.Url)
		if err != nil {
			log.Printf("error fetching url %s: %v", r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _,item := range parserResult.Items {
			log.Printf("Got item: %v: ", item)
		}
	}
}