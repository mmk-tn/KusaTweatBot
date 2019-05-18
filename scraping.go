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

	svgHTML, _ := doc.Find(".js-calendar-graph").Html()

	return svgHTML
}
