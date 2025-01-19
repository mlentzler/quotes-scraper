package main

import (
	"fmt"

	"github.com/mlentzler/quotes_scraper/internal/crawler"
	"github.com/mlentzler/quotes_scraper/internal/parser"
)

func main() {
	url := "https://quotes.toscrape.com"
	fmt.Printf("Fetching URL: %s\n", url)

	html, err := crawler.FetchAllPages(url)
	if err != nil {
		fmt.Errorf("Error Occured while fetching: %v\n", err)
		return
	}

	qoutes, err := parser.ParseQuotes(html)
	if err != nil {
		fmt.Errorf("Error occured while parsing: v%\n", err)
		return
	}
	for _, qoute := range qoutes {
		fmt.Printf("Quote: %s\n", qoute.Text)
		fmt.Printf("Author: %s\n", qoute.Author)
	}
}
