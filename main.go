package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel
	// we expect to share values of type, string
	c := make(chan string)

	// go routine
	for _, link := range links {
		go checkLink(link, c)
	}

	// main routine listening for a response from a channel
	// blocking call -main routine won't finish w/o
	// a child routine finishing first

	// for {} is an inifinite loop
	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
