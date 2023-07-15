package main

import "fmt"

func (d data) printEntries() {
	for i, entry := range d.storedUrls {
		fmt.Println("Entry: ", i, "Url: ", entry.shortenedUrl)
	}
}
