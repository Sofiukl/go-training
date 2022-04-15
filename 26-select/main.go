package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	stop := make(chan string)
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			if i == 8 {
				stop <- "completed"
				break
			}
		}
	}()

	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	// for i := 0; i < 2; i++ {
	for {
		go func() {
			select {
			case task := <-ch:
				go doTask(task)
				return
			case msg1 := <-stop:
				fmt.Println(msg1)
				break
			}
		}()
	}
	//}
}

func doTask(task int) {
	//time.Sleep(1000)
	link := "http://facebook.com"
	if _, err := http.Get(link); err != nil {
		remarks := fmt.Sprintf("Looks like %s is down", link)
		fmt.Println(time.Now().String(), remarks)
		return
	}
	remarks := fmt.Sprintf("Looks like %s is up", link)
	fmt.Println(time.Now().String(), remarks)
	fmt.Println("Processing task ", task)
}
