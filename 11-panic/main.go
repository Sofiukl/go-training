package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println("1")
	fmt.Println("test")
	f1()
}

/* The panic() function in Go Language is similar to exceptions raised at runtime when an error is encountered.
panic() is either raised by the program itself when an unexpected error occurs
or the programmer throws the exception on purpose for handling particular errors. */

func f1() {
	//defer fmt.Println("2")
	//defer fmt.Println("end")

	colors := make(map[string]int)
	colors["red"] = 1
	colors["yellow"] = 2

	panic("oops! error")
}
