package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	name string
	age  int
}

func main() {
	e1 := employee{"sofikul", 23}
	e2 := employee{"sofikul", 23}

	fmt.Println(reflect.DeepEqual(e1, e2))

	amap := map[string]int{
		"white": 1,
		"red":   2,
	}

	anotherMap := map[string]int{
		"white": 1,
		"red":   3,
	}

	fmt.Println(reflect.DeepEqual(amap, anotherMap))

}
