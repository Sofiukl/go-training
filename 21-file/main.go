package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("/Users/optlptp296/Documents/elk/logs/spring-boot-elk.log")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
