go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/net/html"
)

type Crawler struct {
	urls    chan string
	visited map[string]bool
	mu      sync.Mutex
}

func NewCrawler() *Crawler {
	return &Crawler{
		urls:    make(chan string),
		visited: make(map[string]bool),
	}
}

func (c *Crawler) Start(seedURL string) {
	go func() {
		c.urls <- seedURL
	}()

	for i := 0; i < 5; i++ {
		go c.crawl()
	}
}

func (c *Crawler) crawl() {
	for url := range c.urls {
		c.mu.Lock()
		if c.visited[url] {
			c.mu.Unlock()
			continue
		}
		c.visited[url] = true
		c.mu.Unlock()

		fmt.Println("Crawling", url)

		resp, err := http.Get(url)
		if err!= nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err!= nil {
			log.Println(err)
			continue
		}

		c.extractLinks(doc, url)
	}
}

func (c *Crawler) extractLinks(doc *html.Node, base string) {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, attr := range doc.Attr {
			if attr.Key == "href" {
				link, err := url.Parse(attr.Val)
				if err!= nil {
					log.Println(err)
					continue
				}

				link = base.ResolveReference(link)

				c.mu.Lock()
				if!c.visited[link.String()] {
					c.urls <- link.String()
				}
				c.mu.Unlock()
			}
		}
	}

	for child := doc.FirstChild; child!= nil; child = child.NextSibling {
		c.extractLinks(child, base)
	}
}

func main() {
	c := NewCrawler()
	c.Start("https://example.com")

	select {} 
}
