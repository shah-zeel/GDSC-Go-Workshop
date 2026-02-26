package hn

// Post represents a Hacker News item scraped from the front page.
type Post struct {
	Rank     int
	Title    string
	URL      string
	Points   int
	Comments int
	ID       string // the HN item id (if present)
}