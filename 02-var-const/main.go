package main

import "fmt"

// global constants
const (
	i = 5
	j = 10
)

const k = 5

func main() {
	const l = 20
	// const b = 30

	a := l * k
	fmt.Println(a)
}
