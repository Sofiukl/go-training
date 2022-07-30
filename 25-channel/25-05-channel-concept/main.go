package main

import "fmt"

func main() {
	ch := make(chan int)

	go testFn(ch)
	fmt.Println(<-ch) // blocking
}

func testFn(ch chan int) {
	ch <- 10
}
