package main

import (
	"log"
	"net/http"
	"net/url"
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func Crawl(url string, depth int, m *message) {
	// TODO: adapt to depth
	defer func() {
		m.quit <- 0
	}()
	list, _:= Fetch(url)
	list = append([][]string{{"title", "subtitle", "description", "url"}}, list...)
	WriteCSV(list)
}

func Fetch(u string) (list [][]string, err error) {
	baseUrl, err := url.Parse(u)
	if err != nil {
		return
	}

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// Selectors to fetch in html
	doc.Find(".s-mh").Each(func(_ int, srg *goquery.Selection) {
		srg.Find(".s-repeatable-item").Each(func(_ int, s *goquery.Selection) {
			title := s.Find(".s-font-heading").Text()
			subtitle := s.Find(".s-item-subtitle").Text()
			description := s.Find(".s-item-text").Text()
			url, _ := s.Find(".s-component-content > a").Attr("href")
			list = append(list, []string{title, subtitle, description, url})
		})
	})

	return
}

func WriteCSV(list [][]string) {
	file, err := os.Create("./output.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.WriteAll(list)
	writer.Flush()
}

func main() {
	m := newMessage()
	go m.execute()
	fetchURL := "http://hoge.com"
	m.req <- &request{
		url:   fetchURL,
		depth: 1,
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
