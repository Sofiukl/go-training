package main

import (
	"fmt"
	"net/http"
	"time"
)

// Main Routine
func main() {
	ch := make(chan string)
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
	}

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch { // blocking until get a value from channel
		// annonymous function, immediately invoked
		go func(l string) {
			time.Sleep(5000 * time.Millisecond)
			checkLink(l, ch)
		}(l) // similar like javascript IIFE
	}

}

func checkLink(link string, ch chan string) {
	if _, err := http.Get(link); err != nil {
		remarks := fmt.Sprintf("Looks like %s is down", link)
		fmt.Println(time.Now().String(), remarks)
		ch <- link
		return
	}
	remarks := fmt.Sprintf("Looks like %s is up", link)
	fmt.Println(time.Now().String(), remarks)
	ch <- link
}
