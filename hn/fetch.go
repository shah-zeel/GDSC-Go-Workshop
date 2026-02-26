package hn

// TODO: IMPORT the required packages:
// - net/http
// - github.com/PuerkitoBio/goquery
// - fmt (optional)

const HN_URL = "https://news.ycombinator.com/"

// FetchDocument should:
// 1. Send an HTTP GET request to HN_URL
// 2. Check for errors
// 3. Check that status code == 200
// 4. Convert the response body into a goquery.Document
// 5. Return the document
//
// HINTS:
// - Use http.Get()
// - Always defer res.Body.Close()
// - Use goquery.NewDocumentFromReader(res.Body)
//
// Signature is given. Students must implement it.
func FetchDocument() (/* what type should this return? */, error) {

	// TODO: Implement HTTP request

	// TODO: Convert response body to goquery document

	// TODO: Return the document

	return nil, nil
}