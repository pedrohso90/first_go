package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	startTime := time.Now()

	// Model 1
	if len(os.Args) > 1 {
		fmt.Println("array: ", os.Args[1:])
		fmt.Println("string: ", strings.Join(os.Args[1:], ", "))
	} else {
		fmt.Println("cannot be empty")
		fmt.Println(time.Since(startTime).Seconds(), "Time execution for{} method")
		os.Exit(1)
	}

	// Model 2
	if len(os.Args) <= 1 {
		fmt.Println("ex.: ", os.Args[0], "<some arguments>")
		fmt.Println(time.Since(startTime).Seconds(), "Time execution for{} method")
		os.Exit(1)
	}

	fmt.Println("array: ", os.Args[1:])
	fmt.Println("string: ", strings.Join(os.Args[1:], ", "))

	fmt.Println(time.Since(startTime).Seconds(), "Time execution for{} method")
}
