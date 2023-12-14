package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func makeRequest() error {
	// manages the network communication and is responsible for the execution of the attached callbacks while a collector job is running. To work with colly, you have to initialize a Collector
	c := colly.NewCollector()

	// Request starts a collector job by creating a custom HTTP request
	method := "GET"
	URL := "cats.com"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	err := c.Request(method, URL, nil, nil, nil)

	return err
}
