package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	lists := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.linkedin.com/",
		"https://github.com/",
		"https://slack.com/",
	}
	c := make(chan string)
	for _, list := range lists {
		go checkWebsite(list, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkWebsite(link, c)
		}(l)
	}
}

func checkWebsite(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, " might be down!")
		c <- url
		return
	}

	fmt.Println(url, " is Up.")
	c <- url
}
