package main

import (
	"fmt"
	"os"
)

func main(){
	p := os.Args[1:]
	fmt.Println("ola", p)
}
