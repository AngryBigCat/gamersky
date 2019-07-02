package gamersky

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/AngryBigCat/gamersky/models"
	"github.com/AngryBigCat/gamersky/utils"

	"github.com/AngryBigCat/gamersky/engine"
	"github.com/PuerkitoBio/goquery"
)

type NewsList struct {
	Status     string  `json:"status"`
	TotalPages float64 `json:"totalPages"`
	Body       string  `json:"body"`
}

const listRe = `^jQuery.*?\((.*)\);$`

func ParseNewsList(content []byte) engine.ParserResult {
	newsList := ParserNewsListToType(content)

	reader := strings.NewReader(newsList.Body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	var parserResult engine.ParserResult
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		subject := s.Find(".dh").Text()
		title := s.Find(".tt").Text()
		href, _ := s.Find(".tt").Attr("href")
		desc := s.Find(".con .txt").Text()
		img, _ := s.Find("img").Attr("src")
		datetime := utils.DatetimeToUnix(s.Find(".con .tem .time").Text())

		parserResult.Requests = append(parserResult.Requests, engine.Request{
			Url:        href,
			ParserFunc: engine.NilParserFunc,
		})

		new := models.New{
			Subject:     subject,
			Title:       title,
			Href:        href,
			Image:       img,
			Description: desc,
			PublishAt:   datetime,
		}

		parserResult.Items = append(parserResult.Items, &new)

		fmt.Printf("Review %d: %s - %s - %s - %s - %d - %s \n", i, subject, title, href, desc, datetime, img)

		DB.Create(&new)
	})

	/* 正则
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
	*/

	return parserResult
}

func ParserNewsListToType(content []byte) NewsList {
	re := regexp.MustCompile(listRe)
	submatch := re.FindSubmatch(content)

	var newsList NewsList
	err := json.Unmarshal(submatch[1], &newsList)
	if err != nil {
		log.Println(err)
	}
	return newsList
}
