package main

import "fmt"

// example of closure function like javascript
func nextSeqGen() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextSeqFn := nextSeqGen()
	fmt.Println(nextSeqFn())
	fmt.Println(nextSeqFn())
	fmt.Println(nextSeqFn())
	fmt.Println(nextSeqFn())
}
