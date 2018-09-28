package news

import (
	"fmt"
	"gamersky/engine"
	"regexp"
)

const newsListRe = `<li>.*?<a class=\\"dh\\".*?>(.*?)</a>.*?<a class=\\"tt\\" href=\\"(.*?)\\".*?>(.*?)</a>.*?<div class=\\"txt\\">(.*?)</div>.*?<img src=\\"(.*?)\\".*?<div class=\\"time\\">(.*?)</div>.*?</li>`

func ParseNewsList(content []byte) engine.ParserResult {
	compile := regexp.MustCompile(newsListRe)
	submatch := compile.FindAllSubmatch(content, -1)
	parserResult := engine.ParserResult{}
	for k, v := range submatch {
		fmt.Printf("%d: %s  %s\n", k, v[3], v[2])
		parserResult.Items = append(parserResult.Items, v[3])
		parserResult.Requests = append(parserResult.Requests, engine.Request{
			Url:        string(v[2]),
			ParserFunc: engine.NilParserFunc,
		})
	}
	return parserResult
}
