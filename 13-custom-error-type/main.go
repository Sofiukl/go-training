package main

import (
	"fmt"
	"os"
)

/*
Go provides below interface for error handling
type error interface {
   Error() string
}

Any type which implements this contract can be considered of type error.
Check below code to check how to create custom error

*/

type invalidArgument struct {
	arg string
}

func (i invalidArgument) Error() string {
	return fmt.Sprintf("Argument %s is invalid", i.arg)
}

func main() {
	arg1 := os.Args[1] // accessing command line arguments
	if _, err := parseArgument(arg1); err != nil {
		fmt.Println(err)
	}
}

// func with multiple return type like python
func parseArgument(arg string) (string, error) {
	return "", invalidArgument{arg: arg}
}
