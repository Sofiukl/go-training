package main

import (
	"fmt"
	"reflect"
)

func main() {

	// static type
	var i int
	fmt.Println(i)

	// dynamic type
	var j = 10
	fmt.Println(j)

	// short form of dynamic type
	k := 12
	fmt.Println(k)

	// mixed type
	var a, b, c = 3, 4, "foo"
	fmt.Println(a, b, c)

	// check what type infered by go
	fmt.Println(reflect.TypeOf(j))
}
