package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	p := os.Args[1:]
	s := strings.Join(p,"")
	fmt.Println("array: ", p)
	fmt.Println("string: ", s)
}
