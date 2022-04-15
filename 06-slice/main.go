package main

import "fmt"

func main() {
	/*
		Arrays have their place, but they’re a bit inflexible, so you don’t see them too often in Go code.
		Slices, though, are everywhere. They build on arrays to provide great power and convenience.
		The type specification for a slice is []T, where T is the type of the elements of the slice.
		Unlike an array type, a slice type has no specified length.
		A slice literal is declared just like an array literal, except you leave out the element count:
	*/

	var mySlice []string

	mySlice = append(mySlice, "test string")
	mySlice = append(mySlice, "test string1")
	mySlice = append(mySlice, "test string2")

	for index, val := range mySlice {
		fmt.Println(index, val)
	}
}
