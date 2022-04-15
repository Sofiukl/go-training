package main

import "fmt"

func main() {
	slice1 := []int{1, 5, 7, 2}
	len := len(slice1)
	slice2 := []int{}
	for i := range slice1 {
		slice2 = append(slice2, slice1[len-i-1])
	}

	fmt.Println(slice2)
}
