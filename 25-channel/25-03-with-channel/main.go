package main

import (
	"fmt"
	"net/http"
)

func main() {
	ch := make(chan string)
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
	}

	for _, link := range links {
		go checkLinks(link, ch)
	}

	for {
		fmt.Println(<-ch)
	}

}

func checkLinks(link string, ch chan string) {
	if _, err := http.Get(link); err != nil {
		remarks := fmt.Sprintf("Looks like %s is down", link)
		ch <- remarks
		return
	}
	remarks := fmt.Sprintf("Looks like %s is up", link)
	ch <- remarks
}
