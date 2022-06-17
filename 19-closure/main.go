package main

import "fmt"

// A closure is a function that references variables outside its own function body (scope).
// In other words, a closure is an inner function that has access to the variables in the scope in which it was created.
// This applies even when the outer function finishes execution and the scope gets destroyed.

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
