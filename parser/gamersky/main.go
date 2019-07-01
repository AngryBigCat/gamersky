package gamersky

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AngryBigCat/gamersky/engine"
)

const newsListTotalRe = `.*?"totalPages":(\d+),.*`

func ParserMain(content []byte) engine.ParserResult {
	compile := regexp.MustCompile(newsListTotalRe)
	result := compile.FindSubmatch(content)
	total, _ := strconv.Atoi(string(result[1]))

	parserResult := engine.ParserResult{}
	for i := 1; i <= total; i++ {
		parserResult.Requests = append(parserResult.Requests, engine.Request{
			Url:        "https://db2.gamersky.com/LabelJsonpAjax.aspx?callback=jQuery18303958816852369551_1535642187524&jsondata=%7B%22type%22%3A%22updatenodelabel%22%2C%22isCache%22%3Atrue%2C%22cacheTime%22%3A60%2C%22nodeId%22%3A%2211007%22%2C%22isNodeId%22%3A%22true%22%2C%22page%22%3A" + fmt.Sprintf("%d", i) + "%7D&_=1535642221589",
			ParserFunc: ParseNewsList,
		})
		parserResult.Items = append(parserResult.Items, i)
	}
	return parserResult
}
