package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	feedParser := gofeed.NewParser()

	feed, err := feedParser.ParseURL("https://alexisglez.netlify.app/blogs/feed.json")
	if err != nil {
		log.Fatalf("Error while getting feed: %v", err)
	}

	// Get latests blogs
	var blogs bytes.Buffer
	for i := 0; i < 5; i++ {
		rssItem := feed.Items[i]
		blogs.WriteString("- [" + rssItem.Title + "](" + rssItem.Link + ")\n")
	}

	date := time.Now().Format("1 Jan 2000")
	updated := "Last updated on " + date + "."
	blogs.WriteString(updated + "\n")

	// Read original README
	content, err := ioutil.ReadFile("originalReadme.md")
	if err != nil {
		log.Fatalf("Cannot read original readme: %v", err)
	}

	stringyContent := string(content)

	// Add latest blogs to original readme
	readme := strings.Replace(stringyContent, "<!-- My Blogs go here -->", blogs.String(), 1)

	// Create a new readme file at the root
	file, err := os.Create("../README.md")
	if err != nil {
		log.Fatalf("Cannot create new readme: %v", err)
	}
	defer file.Close()

	_, err = io.WriteString(file, readme)
	if err != nil {
		log.Fatalf("Cannot write content in new readme: %v", err)
	}

	err = file.Sync()
	if err != nil {
		log.Fatalf("Cannot save new readme in storage: %v", err)
	}
}
