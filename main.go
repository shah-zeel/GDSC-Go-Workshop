package main

import (
	"flag"
	"log"

	"hntemplate/hn"
)

// This file ONLY handles CLI flags.
// Should not need to modify this much.
func main() {
	mode := flag.String("mode", "trending", "trending | top10")
	n := flag.Int("n", 10, "number of results")
	flag.Parse()

	switch *mode {
	case "trending":
		if err := hn.Trending(*n); err != nil {
			log.Fatal(err)
		}
	case "top10":
		if err := hn.Top10(*n); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unknown mode")
	}
}