package main

import "fmt"

func main() {
	fmt.Println("Hello, web.")
	links, err := makeRequest("https://80000hours.org")
	if err != nil {
		fmt.Println(err)
	}

	countedDomains := countDomains(links)

	fmt.Println(countedDomains)
}
