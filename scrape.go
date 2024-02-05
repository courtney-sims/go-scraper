package main

import (
	"fmt"
	"strings"
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

func countDomains(links []string) map[string]int {
	// Go learning: The len function returns the number of elements in a slice
	fmt.Println("Number of links found:", len(links))
	uniqueDomains := 0
	countedDomains := make(map[string]int)

	// Cut out the url of the site we're scraping
	// This won't help with relative paths though. We either need to also check for starts with or return absolute paths to the slice
	for _, link := range links {
		// separate by slashes in domain
		separation := strings.Split(link, "/")
		protocol := separation[0]

		// if 0th element is not http or https, it's a relative path to the site, so we'll ignore that
		if protocol != "http:" && protocol != "https:" {
			continue
		}

		// pull domain only once we've found a full link
		domain := separation[2]
		if domain == "80000hours.org" {
			continue
		}
		uniqueDomains++
		fmt.Println(domain)

		// if we've already seen this domain, increment the count
		// otherwise, add it to the map
		count, ok := countedDomains[domain]
		if ok {
			countedDomains[domain] = count + 1
		} else {
			countedDomains[domain] = 1
		}
	}
	fmt.Println("Number of unique domains found:", len(countedDomains))
	return countedDomains
}
