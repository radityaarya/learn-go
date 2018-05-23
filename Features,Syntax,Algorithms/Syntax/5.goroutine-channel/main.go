package main

import (
	"fmt"
	"net/http"
	"time"
)

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

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	ch := make(chan string)
	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(4 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}
