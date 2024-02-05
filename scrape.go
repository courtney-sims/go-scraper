package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func makeRequest(url string) ([]string, error) {
	var links []string
	// Colly learning: The collector manages the network communication and is responsible for the execution of the attached callbacks while
	// a collector job is running. To work with colly, you have to initialize a Collector.
	// Go learning: short declarations are only possible inside a function
	c := colly.NewCollector()

	// Be nice to the sites we crawl
	c.Limit(&colly.LimitRule{
		// Filter domains affected by tis rule
		//DomainGlob: "80000hours.org/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	method := "GET"

	// Printing what's going on
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	//Parsing the html for links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//Go learning: Appending to a slice is a constant time operation
		links = append(links, e.Attr("href"))
	})

	// Request starts a collector job by creating a custom HTTP request.
	err := c.Request(method, url, nil, nil, nil)

	return links, err
}

func sortLinks(links []string) {
	// Go learning: The len function returns the number of elements in a slice
	fmt.Println("Number of links found:", len(links))
}
