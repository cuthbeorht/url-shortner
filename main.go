package main

import "fmt"

type entry struct {
	originalUrl  string
	shortenedUrl string
}

type data struct {
	storedUrls []entry
}

func main() {

	var url string

	fmt.Println("Please enter a URL to shorten:  ")
	fmt.Scanf("%s", &url)

	first := entry{originalUrl: url, shortenedUrl: "https://short.en/url"}

	list := data{[]entry{first}}

	list.printEntries()
}
