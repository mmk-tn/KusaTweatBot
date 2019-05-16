package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func getSvgHTML(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error loading GitHub page.")
		return ""
	}

	svgHTML := ""
	doc.Find(".js-calendar-graph").Each(func(_ int, s *goquery.Selection) {
		html, _ := s.Html()
		svgHTML += html
	})

	return svgHTML
}
