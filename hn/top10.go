package hn

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// atoiDigits extracts the first sequence of digits and converts to int.
func atoiDigits(s string) int {
	re := regexp.MustCompile(`\d+`)
	m := re.FindString(s)
	if m == "" {
		return 0
	}
	v, _ := strconv.Atoi(m)
	return v
}

// normalizeHref makes relative HN links absolute.
func normalizeHref(href string) string {
	if href == "" {
		return ""
	}
	if strings.HasPrefix(href, "item?id=") {
		return hnBaseURL + href
	}
	if strings.HasPrefix(href, "//") {
		return "https:" + href
	}
	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		return href
	}
	return hnBaseURL + href
}

// extractTopPosts scrapes up to `limit` posts from the given document.
func extractTopPosts(doc *goquery.Document, limit int) ([]Post, error) {
	var posts []Post

	doc.Find(".athing").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if len(posts) >= limit {
			return false
		}

		rankText := strings.TrimSpace(s.Find(".rank").Text())
		rank := atoiDigits(rankText)
		if rank == 0 {
			rank = i + 1
		}

		titleSel := s.Find(".titleline a")
		if titleSel.Length() == 0 {
			titleSel = s.Find("a.storylink")
		}
		title := strings.TrimSpace(titleSel.Text())
		href, _ := titleSel.Attr("href")
		href = strings.TrimSpace(href)
		href = normalizeHref(href)

		idAttr, _ := s.Attr("id")

		// find the sibling row that contains the .subtext (points/comments)
		sub := s.Next()
		steps := 0
		for sub.Length() > 0 && sub.Find(".subtext").Length() == 0 && steps < 4 {
			sub = sub.Next()
			steps++
		}

		points := 0
		comments := 0

		if sub.Find(".score").Length() > 0 {
			pointsText := sub.Find(".score").Text()
			points = atoiDigits(pointsText)
		}

		if sub.Length() > 0 {
			anchors := sub.Find("a")
			if anchors.Length() > 0 {
				lastText := strings.TrimSpace(anchors.Last().Text())
				if strings.Contains(lastText, "comment") || strings.Contains(lastText, "discuss") {
					comments = atoiDigits(lastText)
				} else {
					found := false
					anchors.EachWithBreak(func(_ int, a *goquery.Selection) bool {
						txt := strings.TrimSpace(a.Text())
						if strings.Contains(txt, "comment") {
							comments = atoiDigits(txt)
							found = true
							return false
						}
						return true
					})
					if !found {
						comments = 0
					}
				}
			}
		}

		p := Post{
			Rank:     rank,
			Title:    title,
			URL:      href,
			Points:   points,
			Comments: comments,
			ID:       idAttr,
		}
		posts = append(posts, p)
		return true
	})

	if len(posts) > limit {
		posts = posts[:limit]
	}
	return posts, nil
}

// printPosts prints posts in a readable CLI format.
func printPosts(posts []Post) {
	for i, p := range posts {
		fmt.Printf("#%d — %s\n", i+1, p.Title)
		fmt.Printf("    URL: %s\n", p.URL)
		fmt.Printf("    Points: %d | Comments: %d\n", p.Points, p.Comments)
	}
}

// TopPosts is the exported entrypoint that fetches and prints the top N posts.
func TopPosts(limit int) error {
	doc, err := fetchDocument(hnBaseURL)
	if err != nil {
		return err
	}
	posts, err := extractTopPosts(doc, limit)
	if err != nil {
		return err
	}
	if len(posts) == 0 {
		return errors.New("no posts found; page structure may have changed")
	}

	// sort by Rank for sanity
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Rank < posts[j].Rank
	})

	for i := range posts {
		if posts[i].Rank == 0 {
			posts[i].Rank = i + 1
		}
	}

	printPosts(posts)
	return nil
}