package main

import (
	"fmt"
)

type Node struct {
	name  string
	email string
	adj   []Node
}

func main() {

	var names [][]string

	names = append(names, []string{"sofikul", "nusu"}, []string{"sofikul", "nusu"})
	fmt.Println(names)

}
