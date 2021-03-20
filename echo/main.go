package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	// read arguments
	p := os.Args[1:]
	// array to string
    s := strings.Join(os.Args[1:], ", ")
	// check if argument is empty
	if len(os.Args) > 1 {
		fmt.Println("array: ", p)
		fmt.Println("string: ", s)
	} else {
		fmt.Println("cannot be empty")
		os.Exit(1)
	}
}
