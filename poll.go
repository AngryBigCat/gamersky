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

	for {
		listPage := 1
		url := "https://db2.gamersky.com/LabelJsonpAjax.aspx?callback=jQuery18303958816852369551_1535642187524&jsondata=%7B%22type%22%3A%22updatenodelabel%22%2C%22isCache%22%3Atrue%2C%22cacheTime%22%3A60%2C%22nodeId%22%3A%2211007%22%2C%22isNodeId%22%3A%22true%22%2C%22page%22%3A"+ fmt.Sprintf("%d", listPage) +"%7D&_=1535642221589"
		request <- url
		for {
			if <-hasNext {
				listPage++
				url = "https://db2.gamersky.com/LabelJsonpAjax.aspx?callback=jQuery18303958816852369551_1535642187524&jsondata=%7B%22type%22%3A%22updatenodelabel%22%2C%22isCache%22%3Atrue%2C%22cacheTime%22%3A60%2C%22nodeId%22%3A%2211007%22%2C%22isNodeId%22%3A%22true%22%2C%22page%22%3A"+ fmt.Sprintf("%d", listPage) +"%7D&_=1535642221589"
				fmt.Println(url)
				request <- url
				log.Println("下一页有新文章")
			} else {
				listPage = 1
				log.Printf("新闻列表更新了条")
				break
			}
		}

		//now := time.Now().Format("2006-01-02 15:04:05")
		time.Sleep(time.Minute * 60)
	}
}