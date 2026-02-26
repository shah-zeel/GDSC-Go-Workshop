package main

import (
	"flag"
	"log"

	"hntool/hn"
)

func main() {
	mode := flag.String("mode", "trending", "mode: 'trending' or 'top10'")
	top := flag.Int("n", 10, "for 'trending' = number of top keywords")
	minlen := flag.Int("minlen", 3, "minimum token length for trending words")
	limit := flag.Int("limit", 10, "how many top posts to fetch for 'top10' (default 10)")
	flag.Parse()

	switch *mode {
	case "trending":
		if err := hn.Trending(*top, *minlen); err != nil {
			log.Fatalf("trending error: %v", err)
		}
	case "top10":
		if err := hn.TopPosts(*limit); err != nil {
			log.Fatalf("top10 error: %v", err)
		}
	default:
		log.Fatalf("unknown mode: %s. allowed: trending, top10", *mode)
	}
}