package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
	}

	for _, link := range links {
		checkLinks(link) // sync call
	}
}

func checkLinks(link string) {
	if _, err := http.Get(link); err != nil { // familiar with this composit sytax, instead of null here it is nil
		remarks := fmt.Sprintf("Looks like %s is down", link)
		fmt.Println(time.Now().String(), remarks)
		return
	}
	remarks := fmt.Sprintf("Looks like %s is up", link)
	fmt.Println(time.Now().String(), remarks)
}
