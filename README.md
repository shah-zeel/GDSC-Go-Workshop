# 📰 Hacker News Scraper — Workshop Template

## 🚀 What You Are Building

You are building a web scraper in Go using **goquery**.

Your program will support two modes:

### 1️⃣ Trending Mode

* Scrape all titles on the HN front page.
* Count word frequency.
* Print top N trending keywords.

### 2️⃣ Top10 Mode

* Scrape the first N posts.
* Extract:

  * Rank
  * Title
  * URL
  * Points
  * Comments
* Print formatted output.

---

# 🧰 Requirements

* Go 1.20+

Check Go:

```
go version
```

If Go is not installed:

* Download from: https://go.dev/dl/
* Follow installation instructions for your OS.

# Start a Go project from scratch

1. Create a project folder and go into it:
   ```bash
   mkdir my-go-app && cd my-go-app
   ```
2. Initialize a Go module (replace with your own module path, e.g. `github.com/username/my-go-app`):
   ```bash
   go mod init github.com/username/my-go-app
   ```
3. Create your first file, e.g. `main.go`, and add a simple `package main` with a `main()` function.
4. Run the project:
   ```bash
   go run main.go
   ```

---

# ⚙️ Setup

1. Clone this repo
```bash
git clone https://github.com/shah-zeel/GDSC-Go-Workshop.git
```
2. Navigate into the project folder.
3. Checkout to the `starter-code` branch
```bash
git checkout starter-code
```
4. Install dependencies:

```
go mod tidy
```

This downloads `goquery`.

---

# ▶️ How to Run

### Trending mode:

```
go run . -mode=trending -n=10
```

### Top 10 mode:

```
go run . -mode=top10 -n=10
```

---

# 🧠 What You Need To Implement

All logic is intentionally empty.

You must implement:

---

## Step 1 — HTTP Request (fetch.go)

In `FetchDocument()`:

* Use `http.Get(HN_URL)`
* Check error
* Defer `res.Body.Close()`
* Ensure `res.StatusCode == 200`
* Create goquery document:

  ```
  goquery.NewDocumentFromReader(res.Body)
  ```

Return the document.

---

## Step 2 — Extract Titles (Trending Mode)

In `Trending()`:

1. Call `FetchDocument()`
2. Use:

   ```
   doc.Find(".athing")
   ```
3. Inside each, find:

   ```
   .titleline a
   ```
4. Extract `.Text()`

Store titles in a slice.

---

## Step 3 — Word Processing

* Convert to lowercase
* Remove punctuation using `regexp`
* Split with `strings.Fields`
* Count using:

  ```
  map[string]int
  ```

---

## Step 4 — Sorting

* Convert map into slice of struct:

  ```
  type kv struct { Key string; Value int }
  ```
* Use `sort.Slice()` to sort descending.

---

## Step 5 — Top10 Extraction

Inside `Top10()`:

For each `.athing`:

* Rank → `.rank`
* Title → `.titleline a`
* URL → `.Attr("href")`
* Points → find `.score`
* Comments → last link in subtext row

⚠️ Points/comments are in the NEXT sibling row.

Use:

```
s.Next()
```

---

# 🧩 Hints

Open Hacker News in browser.

Right-click → Inspect → Look at HTML.

You will see:

```
<tr class="athing">
```

Use that selector.

