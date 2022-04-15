package main

import "fmt"

func main() {
	ch := make(chan int)

	go testFn(ch)
	fmt.Println(<-ch) // blocking

	// for i := 0; i < 10; i++ {
	// 	go testFn(ch, i)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }
}

func testFn(ch chan int) {
	ch <- 10
}

// func testFn(ch chan int, i int) {
// 	ch <- i
// }
