package main

import (
	"fmt"

	"github.com/mlentzler/quotes_scraper/internal/crawler"
)

func main() {
  url := "https://quotes.toscrape.com"
  fmt.Printf("Fetching URL: %s\n", url)


  html, err := crawler.Fetch(url)
  if err != nil {
    fmt.Errorf("Error Occured while fetching: %v\n", err)
    return
  }
  
  fmt.Println("Fetched HTML:")
  fmt.Println(html)
}
