package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func Crawl(link string) {
	fmt.Printf("Crawling URL: %s\n", link)
}

func Fetch(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("An Error Occured while fetching: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected Status code: %d", response.StatusCode)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("An Error Occured while reading the body from URL %s: %v", url, err)
	}

	return string(bodyBytes), nil
}

func FetchAllPages(startUrl string) (string, error) {
	currentURL := startUrl

	var builder strings.Builder

	for currentURL != "" {
		html, err := Fetch(currentURL)
		if err != nil {
			fmt.Printf("Error fetching page: %v\n", err)
			break
		}

		builder.WriteString(html)

		nextPage, err := FindNextPage(html)
		if err != nil {
			fmt.Println("No more pages found.")
			break
		}

		currentURL = nextPage
	}

	return builder.String(), nil
}

func FindNextPage(html string) (string, error) {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	// Suche nach `li.next > a`
	nextLink, exists := doc.Find("li.next > a").Attr("href")
	if !exists {
		return "", fmt.Errorf("no next page found")
	}

	// Debugging: Gefundener Link
	fmt.Printf("Debug: Found next link: %s\n", nextLink)

	// Relative URL in absolute URL umwandeln
	if !strings.HasPrefix(nextLink, "http") {
		baseURL := "http://quotes.toscrape.com" // Basis-URL
		nextLink = baseURL + nextLink
	}

	return nextLink, nil
}
