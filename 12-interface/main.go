package main

import "fmt"

// kind of super type
type shape interface {
	draw() // all the contracts are defined here
}

// kind of sub type
type circle struct{}
type rectangle struct{}

func (c circle) draw() {
	fmt.Println("circle")
}

func (c rectangle) draw() {
	fmt.Println("rectangle")
}

/**
	Below describes a similar concept like
	OOP polymorphisom
**/
func main() {
	c1 := circle{}
	r1 := rectangle{}

	shapes := []shape{c1, r1} // slice

	for _, shape := range shapes {
		shape.draw() // at run time it calls the draw function of actual type, polymorphisom
	}
}
