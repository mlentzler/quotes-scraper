package crawler

import (
	"fmt"
	"io"
	"net/http"
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
