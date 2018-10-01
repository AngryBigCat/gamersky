package gamersky

import (
	"gamersky/engine"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"gamersky/utils"
	"gamersky/fetcher"
	"gamersky/models"
	"fmt"
)

var pages string
func ParserNews(content []byte, id int) engine.ParserResult {
	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	utils.LogFatal(err)

	page := doc.Find(".page_css").Find("a")
	if page.Last().Text() == "下一页" {
		nextHref, _ := page.Last().Attr("href")
		content, _ := fetcher.Get(nextHref)
		ParserNews(content, id)
	}

	var count string
	doc.Find(".Mid2L_con").Find("p").Each(func(i int, s *goquery.Selection) {
		html, _:= s.Html()
		count += "<p>" + html + "</p>"
	})
	pages = count + pages

	if page.First().Text() != "上一页" {
		update := models.DB.Table("news").Where("id = ?", id).Update(models.News{
			Content: pages,
		})
		fmt.Println(update)
	}

	return engine.ParserResult{}
}

