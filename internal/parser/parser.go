package parser

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Quote struct {
	Text   string
	Author string
}

func Parse(link string) {
	fmt.Printf("Prasing URL: %s\n", link)
}

func ParseQuotes(html string) ([]Quote, error) {
	reader := strings.NewReader(html)

	var qoutes []Quote

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("Had Trouble parsing the HTML: %v\n", err)
	}
	doc.Find(".quote").Each(func(index int, element *goquery.Selection) {
		qouteText := element.Find(".text").Text()

		author := element.Find(".author").Text()

		qoute := Quote{
			Text:   qouteText,
			Author: author,
		}

		qoutes = append(qoutes, qoute)
	})

	return qoutes, nil
}

func FindNextPage(html string) (string, error) {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	nextLink, exists := doc.Find("a[rel='next']").Attr("href")
	if !exists {
		return "", fmt.Errorf("no next page found")
	}

	return nextLink, nil
}
