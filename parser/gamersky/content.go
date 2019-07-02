package gamersky

import (
	"bytes"

	"github.com/AngryBigCat/gamersky/fetcher"

	"github.com/AngryBigCat/gamersky/engine"
	"github.com/PuerkitoBio/goquery"
)

func ParserNews(content []byte, id int) engine.ParserResult {
	return engine.ParserResult{}
}

func ParserNewsContent(content []byte) string {
	var pages string

	reader := bytes.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(reader)

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
					count += "<p><img src='" + src + "'/><br/>" + s.Text() + "</p>"
				} else {
					count += "<p><img src='" + src + "'/></p>"
				}
			} else {
				if s.Text() != "" {
					count += "<p>" + s.Text() + "</p>"
				}
			}
		}
	})

	return count + pages
}
