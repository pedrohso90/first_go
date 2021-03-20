package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println("array: ", os.Args[1:])
	fmt.Println("string: ", strings.Join(os.Args[1:], ", "))
}
