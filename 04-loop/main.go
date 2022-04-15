package main

import "fmt"

func main() {

	// style 1
	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// }

	i := 0

	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// style 2, style like while loop of other programing lang
	for {
		if i == 12 {
			break
		}
		fmt.Println(i)
		i++
	}
}
