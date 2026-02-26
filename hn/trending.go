package hn

// TODO: Import required packages:
// - fmt
// - strings
// - regexp
// - sort

// Trending should:
// 1. Call FetchDocument()
// 2. Extract all post titles from Hacker News front page
// 3. Split titles into words
// 4. Count word frequency using a map[string]int
// 5. Sort words by frequency
// 6. Print the top N words
//
// IMPORTANT:
// Inspect Hacker News HTML in your browser.
// Each story is inside an element with class "athing".
// The title link is usually inside ".titleline a".
//
// HINT:
// Use doc.Find("CSS_SELECTOR").Each(...)
func Trending(n int) error {

	// TODO: Call FetchDocument()

	// TODO: Extract titles using goquery selectors

	// TODO: Clean words (lowercase, remove punctuation)

	// TODO: Count frequency in a map

	// TODO: Convert map to slice for sorting

	// TODO: Sort slice by frequency (descending)

	// TODO: Print top n words

	return nil
}