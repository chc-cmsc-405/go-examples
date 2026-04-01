package main

import (
	"fmt"
	"net/http"
)

func checkURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "DOWN"
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return "UP"
	}
	return "DOWN"
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://go.dev",
		"http://this-does-not-exist-abc123.com",
	}

	for _, url := range urls {
		status := checkURL(url)
		fmt.Printf("%-45s %s\n", url, status)
	}
}
