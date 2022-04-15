package main

import (
	"fmt"
	"sync"
)

func process(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i, "doing the task")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All have completed the process")
}
