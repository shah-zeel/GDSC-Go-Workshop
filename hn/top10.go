package hn

// TODO: Import required packages:
// - fmt
// - strings
// - strconv
// - regexp

// Top10 should:
// 1. Call FetchDocument()
// 2. Find the first N posts on the page
// 3. Extract:
//      - Rank
//      - Title
//      - URL
//      - Points
//      - Comments
// 4. Print them nicely
//
// HINTS:
// - Each post has class "athing"
// - Rank is in ".rank"
// - Title is in ".titleline a"
// - Points are in ".score"
// - Comments are usually the last link in the subtext row
//
// You may need to navigate to the NEXT sibling element to get points/comments.
func Top10(n int) error {

	// TODO: Call FetchDocument()

	// TODO: Loop through first n ".athing" elements

	// TODO: Extract rank (convert string to int)

	// TODO: Extract title and href

	// TODO: Normalize relative URLs

	// TODO: Move to next sibling row for points/comments

	// TODO: Convert "123 points" -> 123 (int)

	// TODO: Print results

	return nil
}