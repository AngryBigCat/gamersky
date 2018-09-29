package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	"encoding/json"
	"log"
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

type NewsList struct {
	Status string `json:"status"`
	TotalPages float64 `json:"totalPages"`
	Body string `json:"body"`
}

func main() {
	resp, err := http.Get("https://db2.gamersky.com/LabelJsonpAjax.aspx?callback=jQuery183002571377415944931_1538233140968&jsondata=%7B%22type%22%3A%22updatenodelabel%22%2C%22isCache%22%3Atrue%2C%22cacheTime%22%3A60%2C%22nodeId%22%3A%2211007%22%2C%22isNodeId%22%3A%22true%22%2C%22page%22%3A2%7D&_=1538233147845")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bcontent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//
	//bcontent = bytes.Replace(bcontent, []byte("\\r"), []byte(""), -1)
	//bcontent = bytes.Replace(bcontent, []byte("\\n"), []byte(""), -1)
	//bcontent = bytes.Replace(bcontent, []byte("\\t"), []byte(""), -1)
	//bcontent = bytes.Replace(bcontent, []byte("\\"), []byte(""), -1)


	re := regexp.MustCompile(`^jQuery.*?\((.*)\);$`)
	submatch := re.FindSubmatch(bcontent)

	var newslist NewsList
	err = json.Unmarshal(submatch[1], &newslist)
	if err != nil {
		log.Println(err)
	}

	reader := bytes.NewReader([]byte(newslist.Body))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find(".dh").Text()
		title := s.Find(".tt").Text()
		href, _ := s.Find(".tt").Attr("href")
		txt := s.Find(".con .txt").Text()
		time := s.Find(".con .tem .time").Text()
		fmt.Printf("Review %d: %s - %s - %s - %s - %s \n", i, band, title, href, txt, time)
	})
}