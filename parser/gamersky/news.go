package gamersky

import (
	"gamersky/engine"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"gamersky/utils"
	"gamersky/fetcher"
	"gamersky/models"
)

func ParserNews(content []byte, id int) engine.ParserResult {
	models.DB.Create(&models.Contents{
		NewsId: id,
		Content: ParserNewsContent(content),
	})
	return engine.ParserResult{}
}


func ParserNewsContent(content []byte) string {
	var pages string

	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	utils.LogFatal(err)

	page := doc.Find(".page_css").Find("a")
	if page.Last().Text() == "下一页" {
		nextHref, _ := page.Last().Attr("href")
		content, _ := fetcher.Get(nextHref)
		pages = ParserNewsContent(content)
	}

	var count string
	doc.Find(".Mid2L_con,.MidL_con,.MidLcon").Find("p").Each(func(i int, s *goquery.Selection) {
		if s.Find("#pe100_page_contentpage").Length() == 0 {
			if src, exists := s.Find("img").Attr("src"); exists {
				if s.Text() != "" {
					count +=  "<p><img src='" + src + "'/><br/>"+ s.Text() +"</p>"
				} else {
					count +=  "<p><img src='" + src + "'/></p>"
				}
			} else {
				if s.Text() != "" {
					count +=  "<p>" + s.Text() + "</p>"
				}
			}
		}
	})

	return count + pages
}
