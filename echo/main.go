package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	p := os.Args[1:]
    s := strings.Join(os.Args[1:], ", ")
	if len(os.Args) > 1 {
		fmt.Println("array: ", p)
		fmt.Println("string: ", s)
	} else {
		fmt.Println("cannot be empty")
		os.Exit(1)
	}
}
