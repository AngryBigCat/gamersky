package gamersky

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/AngryBigCat/gamersky/models/db"

	"github.com/AngryBigCat/gamersky/fetcher"

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

		// DB.Create(&new)

		/*
			p, err := db.Pool.Get()
			if err != nil {
				fmt.Println("pools error")
			}

			p.(*gorm.DB).Create(&new)
			db.Pool.Put(p)
		*/

		db.Instance.Create(&new)

		createParseDetailWorker(href, new.Id)
	})

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

func createParseDetailWorker(href string, id int) {
	body, err := fetcher.Get(href)
	if err == nil {
		detail := ParserNewsContent(body)

		content := models.Content{
			NewsId:  id,
			Content: detail,
		}
		/*
			p, err := db.Pool.Get()
			if err != nil {
				fmt.Println("pools error")
			}

			p.(*gorm.DB).Create(&content)
			db.Pool.Put(p)
		*/

		db.Instance.Create(&content)
	}
}
