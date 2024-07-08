**Web Crawler**
================

A simple web crawler written in Go that starts at a given seed URL and crawls the web by following links.

**Usage**
-----

To run the crawler, save this code to a file named `crawler.go` and run the following command:

```bash 
go run crawler.go
```

This will start the crawler with the default seed URL `https://example.com`. You can specify a different seed URL by passing it as an argument:
**How it Works**
--------------

The crawler uses a channel to receive URLs to crawl and a map to keep track of visited URLs. It starts multiple goroutines to crawl the web in parallel. Each goroutine sends an HTTP request to the URL, parses the HTML response, and extracts links from the page. The extracted links are added to the channel to be crawled.

**Features**
--------

* Crawls the web starting from a given seed URL
* Follows links to crawl multiple pages
* Keeps track of visited URLs to avoid duplicates
* Runs multiple goroutines to crawl the web in parallel

**Limitations**
------------

* Does not handle robots.txt files
* Does not respect crawl delays
* Does not handle different types of links (e.g. JavaScript-generated links)
* May not work correctly for all websites

