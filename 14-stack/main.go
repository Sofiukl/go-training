package main

import "fmt"

// Stack - the Stack type (LIFO)

// func append(s []T, vs ...T) []T
// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

type Stack []int

func (s *Stack) push(v int) {
	// s is an pointer to Stack. *s denotes the value of s, &s denotes the memory address of s
	// here we have used pointer receiver, please refer L8 Receiver
	*s = append((*s), v)
}

func (s *Stack) pop() {
	index := len(*s) - 1
	*s = (*s)[:index]
}

func (s *Stack) print() {
	for _, v := range *s {
		fmt.Println(v)
	}
}

func main() {
	var s Stack
	s.push(1)
	s.push(2)
	s.push(3)

	s.pop()

	s.print()

}
