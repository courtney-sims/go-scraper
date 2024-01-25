package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func makeRequest() error {
	// manages the network communication and is responsible for the execution of the attached callbacks while a collector job is running. To work with colly, you have to initialize a Collector
	// short declarations are only possible inside a function
	c := colly.NewCollector()

	// Request starts a collector job by creating a custom HTTP request
	method := "GET"
	// Get dogs from multiple shelters and add information about all of them in a csv
	// Return the oldest dog
	URL := "https://austinpetsalive.org/adopt/dogs"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	err := c.Request(method, URL, nil, nil, nil)

	return err
}
