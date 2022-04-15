package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var lock = &sync.Mutex{}

type single struct{}

var singletone *single

func getInstance() *single {

	if singletone == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletone == nil {
			fmt.Println("Creating new instance")
			singletone = &single{}
		} else {
			fmt.Println("Instance already present 1")
		}
	} else {
		fmt.Println("Instance already present 2")
	}
	return singletone
}

func getInstanceOnce() *single {

	if singletone == nil {
		once.Do(
			func() {
				fmt.Println("Creating new instance")
				singletone = &single{}
			})

	} else {
		fmt.Println("Instance already present")
	}
	return singletone
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance() // go routine, call concurrently
	}
	fmt.Scanln()
}
