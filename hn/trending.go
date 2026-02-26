package hn

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// simple stopword set; students can extend this file as a challenge
var stopwords = map[string]bool{
	"the": true, "and": true, "a": true, "to": true, "of": true,
	"in": true, "for": true, "on": true, "is": true, "with": true,
	"by": true, "an": true, "from": true, "that": true, "as": true,
	"are": true, "this": true, "be": true, "was": true, "it": true,
	"you": true, "your": true, "have": true, "has": true, "i": true,
}

// extractTitles returns the list of story titles from an HN front page document.
func extractTitles(doc *goquery.Document) []string {
	var titles []string
	doc.Find(".athing").Each(func(i int, s *goquery.Selection) {
		// preferred modern selector
		if t := s.Find(".titleline a").Text(); strings.TrimSpace(t) != "" {
			titles = append(titles, strings.TrimSpace(t))
			return
		}
		// fallback older selector
		if t := s.Find("a.storylink").Text(); strings.TrimSpace(t) != "" {
			titles = append(titles, strings.TrimSpace(t))
			return
		}
	})

	// final fallback
	if len(titles) == 0 {
		doc.Find("a.storylink, .titleline a").Each(func(i int, s *goquery.Selection) {
			t := strings.TrimSpace(s.Text())
			if t != "" {
				titles = append(titles, t)
			}
		})
	}
	return titles
}

// tokenize returns lower-case alphanumeric tokens from input text.
func tokenize(text string) []string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	clean := re.ReplaceAllString(text, " ")
	clean = strings.ToLower(clean)
	words := strings.Fields(clean)
	return words
}

// Trending fetches HN front page titles, counts token frequency, and prints top nTop words.
// Exported so students can call it from tests or other packages.
func Trending(nTop int, minLen int) error {
	doc, err := fetchDocument(hnBaseURL)
	if err != nil {
		return err
	}

	titles := extractTitles(doc)
	if len(titles) == 0 {
		return errors.New("no titles found; page structure may have changed")
	}

	freq := make(map[string]int)
	for _, t := range titles {
		for _, w := range tokenize(t) {
			if len(w) < minLen {
				continue
			}
			if stopwords[w] {
				continue
			}
			freq[w]++
		}
	}

	type kv struct {
		Key string
		Val int
	}
	var pairs []kv
	for k, v := range freq {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Val == pairs[j].Val {
			return pairs[i].Key < pairs[j].Key
		}
		return pairs[i].Val > pairs[j].Val
	})

	if nTop > len(pairs) {
		nTop = len(pairs)
	}
	fmt.Printf("Top %d keywords on Hacker News front page (titles):\n", nTop)
	for i := 0; i < nTop; i++ {
		fmt.Printf("%2d. %s — %d\n", i+1, pairs[i].Key, pairs[i].Val)
	}
	return nil
}