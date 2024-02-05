# go-scraper

A terminal application for scraping websites and analyzing where their links go.

This is the first project in my Year of Go. Comments hold tidbits I learn as I code.

# Learnings

1. When writing html, clear class and div names with good groupings of related elements make web scraping easier. The lack of which makes it harder and hackier. My first plan with this webscraper was to ask the question - what is the oldest adoptable animal from a variety of shelters and species? And how often does this change? As I began parsing the HTML from a particular shelter, I struggled to connect the items I needed to in order to make this happen. So I switched to another plan.

# Resources

- https://pkg.go.dev/github.com/gocolly/colly/v2#Collector.Request
- https://go-colly.org/docs/examples/coursera_courses/
- https://haydz.github.io/2020/04/12/ParsingStrings.html
- https://golang.cafe/blog/how-to-check-if-a-map-contains-a-key-in-go
