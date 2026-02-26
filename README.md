# 📰 Hacker News Scraper & Trending Analyzer (Go + goquery)

## 🚀 What This Project Does

This project is a **Hacker News web scraper built in Go** using the `goquery` library.

It supports two main modes:

### 1️⃣ Trending Topic Analyzer

Fetches the Hacker News front page, extracts all post titles, analyzes the words used, and prints the **top trending keywords** based on frequency.

Example output:

```
Top 10 keywords on Hacker News front page (titles):
 1. ai — 4
 2. openai — 3
 3. startup — 3
 4. rust — 2
```

---

### 2️⃣ Top Posts Scraper

Scrapes the top N posts from Hacker News and extracts:

* Rank
* Title
* URL
* Points
* Number of comments

Example output:

```
#1 — Example Post Title
    URL: https://example.com
    Points: 512 | Comments: 123
```

---

This project is designed for:

* Learning web scraping in Go
* Understanding HTML parsing using CSS selectors
* Practicing CLI flag handling
* Structuring Go projects across multiple files
* Extending functionality as a workshop challenge

---

# 🧰 Requirements & Installation

## 1️⃣ Install Go

Make sure you have **Go 1.20+** installed.

Check your version:

```bash
go version
```

If Go is not installed:

* Download from: [https://go.dev/dl/](https://go.dev/dl/)
* Follow installation instructions for your OS.

---

## 2️⃣ Clone or Download This Project

```bash
git clone https://github.com/shah-zeel/GDSC-Go-Workshop.git
```

Or download the project ZIP and extract it.

---

## 3️⃣ Install Dependencies

This project uses:

* `github.com/PuerkitoBio/goquery`

Install dependencies with:

```bash
go mod tidy
```

This will download and install required packages automatically.

---

# 📁 Project Structure

```
hntool/
│
├── go.mod
├── main.go
│
└── hn/
    ├── fetch.go
    ├── models.go
    ├── trending.go
    ├── top10.go
```

### 📌 File Breakdown

### `main.go`

* Entry point of the application
* Parses CLI flags
* Switches between `trending` and `top10` modes
* Calls exported functions from the `hn` package

---

### `hn/fetch.go`

* Handles HTTP requests
* Creates a reusable HTTP client with timeout
* Fetches and parses HTML into a `goquery.Document`

---

### `hn/models.go`

* Defines the `Post` struct
* Used to represent scraped Hacker News posts

---

### `hn/trending.go`

* Extracts titles from the HN front page
* Tokenizes and cleans words
* Filters stopwords
* Counts word frequency
* Sorts and prints trending keywords

---

### `hn/top10.go`

* Extracts:

  * Rank
  * Title
  * URL
  * Points
  * Comments
* Normalizes URLs
* Prints formatted output

---

# ▶️ How to Run

From the root directory:

---

## 🔥 Run Trending Topic Analyzer

```bash
go run . -mode=trending -n=10 -minlen=3
```

### Flags:

| Flag             | Description                       |
| ---------------- | --------------------------------- |
| `-mode=trending` | Run keyword analyzer              |
| `-n=10`          | Number of top keywords to display |
| `-minlen=3`      | Minimum word length               |

---

## 📰 Run Top Posts Scraper

```bash
go run . -mode=top10 -limit=10
```

### Flags:

| Flag          | Description               |
| ------------- | ------------------------- |
| `-mode=top10` | Run top posts scraper     |
| `-limit=10`   | Number of posts to scrape |

---

## 🛠 Build a Binary (Optional)

Instead of using `go run`, you can build a binary:

```bash
go build -o hntool
```

Then run:

```bash
./hntool -mode=top10
```

---

# 🧠 How It Works (High-Level Overview)

1. The program makes an HTTP GET request to:

   ```
   https://news.ycombinator.com/
   ```

2. It parses the returned HTML using `goquery`.

3. It selects elements using CSS selectors like:

   ```
   .athing
   .titleline a
   .score
   ```

4. It extracts relevant data and processes it depending on the mode.
