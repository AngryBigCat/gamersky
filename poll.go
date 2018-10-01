package main

import (
	"log"
	"time"
	"gamersky/models"
	"gamersky/fetcher"
	"gamersky/parser/gamersky"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"gamersky/utils"
	"fmt"
)

func main() {
	request := make(chan string)
	content := make(chan []byte)
	hasNext := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				r := <- request
				c, err := fetcher.Get(r)
				if err != nil {
					log.Println(err)
					continue
				}
				content <- c
			}
		}()
	}

	var news models.News
	var count = 0
	go func() {
		for {
			models.DB.Order("publish_at desc").Limit(1).Find(&news)
			lastTime := news.PublishAt
			newsList := gamersky.ParserNewsListToType(<-content)
			html := strings.NewReader(newsList.Body)
			doc, err := goquery.NewDocumentFromReader(html)
			utils.LogFatal(err)

			key := 0
			next := false
			doc.Find("li").Each(func(i int, s *goquery.Selection) {
				newsTime := utils.DatetimeToUnix(s.Find(".con .tem .time").Text())
				if  next = newsTime > lastTime; next {
					count++
					key++
					subject := s.Find(".dh").Text()
					title := s.Find(".tt").Text()
					href, _ := s.Find(".tt").Attr("href")
					desc := s.Find(".con .txt").Text()
					img, _ := s.Find("img").Attr("src")
					fmt.Println(key, subject, title, href, newsTime, desc, img)
				}
			})
			hasNext <- next
		}
	}()

	url := `https://db2.gamersky.com/LabelJsonpAjax.aspx?callback=jQuery18303958816852369551_1535642187524&jsondata={"type":"updatenodelabel","nodeId":"11007","page":%d}`
	for {
		listPage := 1
		request <-  fmt.Sprintf(url, listPage)
		for {
			if <-hasNext {
				listPage++
				request <- fmt.Sprintf(url, listPage)
				log.Println("下一页有新文章")
			} else {
				log.Printf("本次新闻列表更新了%d条", count)
				listPage = 1
				count = 0
				break
			}
		}

		//now := time.Now().Format("2006-01-02 15:04:05")
		time.Sleep(time.Minute * 1)
	}
}