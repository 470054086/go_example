package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Request struct {
	url      string
	callback func(request chan Request, response *goquery.Document)
	///response *http.Response
}

func main() {
	reqeust := make(chan Request)
	go start_requests(reqeust)
	for {
		select {
		case data := <-reqeust:
			go func() {
				res := get_response(data.url)
				data.callback(reqeust, res)
			}()

		}
	}

}
func get_response(url string) (response *goquery.Document) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	return doc
}
func start_requests(request chan Request) {
	start_url := "http://jhsjk.people.cn/result"
	request <- Request{start_url, parse}
}
func parse(request chan Request, doc *goquery.Document) {
	if doc != nil {
		doc.Find("body > div.w1000.d2list.clearfix > div.fr > ul > li").Each(func(i int, s *goquery.Selection) {
			url, exists := s.Find("a").Attr("href")
			if exists {
				fmt.Println(url)
				request <- Request{url, parse}
			}
		})
	}

}
func article_content(request chan Request, response *http.Response) {
	fmt.Println("抓取完成~")
}