package main

import "fmt"

// Stack - the Stack type
// LIFO
type Stack []int

func (s *Stack) push(v int) {
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
