package main

import "fmt"

// Queue - the Queue type
type Queue []int

func (q *Queue) push(v int) {
	*q = append((*q), v)
}

func (q *Queue) pop() int {
	len := len(*q)
	*q = (*q)[1:len]
	return (*q)[0]
}

func (q *Queue) print() {
	for _, v := range *q {
		fmt.Println(v)
	}
}

func main() {
	var q Queue
	q.push(1)
	q.push(2)
	q.push(3)

	q.pop()

	q.print()

}
