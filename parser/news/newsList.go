package news

import (
	"regexp"
	"gamersky/engine"
)

const newsListRe = `<li>.*?<a class=\\"dh\\".*?>(.*?)</a>.*?<a class=\\"tt\\" href=\\"(.*?)\\".*?>(.*?)</a>.*?<div class=\\"txt\\">(.*?)</div>.*?<img src=\\"(.*?)\\".*?<div class=\\"time\\">(.*?)</div>.*?</li>`
func ParseNewsList(content []byte) engine.ParserResult {
	compile := regexp.MustCompile(newsListRe)
	submatch := compile.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, v := range submatch {
		result.Items = append(result.Items, v[3])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(v[2]),
			ParserFunc: engine.NilParserFunc,
		})
	}
	return result
}