package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	blogTitles, err := GetLAtestBlogTitles("https://golangcode.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Blog titles :")
	fmt.Println(blogTitles)
}

func GetLAtestBlogTitles(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	titles := ""
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += "-" + s.Text() + "\n"
	})
	return titles, nil
}
