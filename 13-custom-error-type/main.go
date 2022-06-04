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

For Java, we do the below to create custom exception

public class InvalidArgument extends Exception { 
    public InvalidArgument(String errorMessage) {
        super(errorMessage);
    }
}

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

/*
In Go it’s idiomatic to communicate errors via an explicit,
separate return value. This contrasts with the exceptions used
in languages like Java and Ruby. Go’s approach makes it easy
to see which functions return errors and to handle them
using the same language constructs employed for any other, non-error tasks.
*/
