# SearcherGO — a lightweight email search engine

SearcherGO is a utility for bulk email extraction using HTML files and saved pages from search engines (Google, Yandex, and any other).

## 🎯 Why is this necessary?

Searching for contacts on websites is a pain:
- **Search engines** — ads, captchas, limitations
- **Corporate websites** — outdated or unreliable information
- **Calls** — endless questions, waiting on hold, unavailable contact

**Email solves everything.** One well-written email gets a response faster than an hour of phone calls.

But clicking every link, searching for emails, and repeating the process is time-consuming and tedious. SearcherGO automates this process: you upload an HTML file with links, and the program finds all emails in seconds.

## 🚀 Quick Start

### Requirements
- Go 1.21+ (for building)

### Installation and Run

```bash
# Clone the repository
git clone https://github.com/NikZha/searcherGo
cd searcherGo

# Build the binary
go build -o searcher.exe

# Run
./searcher.exe
```
Or simply double-click start.bat – the program will open a console with diagnostics and automatically close it when it finishes.

After launching, open a browser and navigate to the address written in the console.

## 📦 What's inside?

| Before (Java/Spring) | After (Go) |
-------------- | ------------- |
| 50+ MB binary + JVM | 10 MB single file |
| Saving to a database (unrequested feature) | Removed - simplicity is more important |
| Search via URL requests (not working due to captchas) | Abandoned this approach |
| Spring Boot server | Starts in milliseconds|

## 💡 How does it work?

1. You save an HTML file with search results (Google, Yandex, or any other)

2. Upload it to the web interface

3. JS collects all links of the form ```<a href="...">```

4. The Go backend processes each link and extracts the email

5. The result is returned as a table

## 📄 License
MIT License - free to use, modify, and redistribute.

## 🙌 Acknowledgments
Inspired by the real need to speed up the search for components and business contacts.