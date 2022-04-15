package main

import "fmt"

func main() {
	age := 3
	if age > 60 {
		fmt.Println("senior")
	} else if age < 18 {
		fmt.Println("minor")
	} else {
		fmt.Println("junior")
	}
}
