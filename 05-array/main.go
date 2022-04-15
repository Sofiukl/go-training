package main

import "fmt"

func main() {
	var arr [6]bool // fixed in size, like java
	arr[0] = true
	arr[3] = true
	fmt.Println(arr)
}
