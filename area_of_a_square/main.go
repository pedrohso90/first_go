package main

import "fmt"

// function area
func area(l int) int {
    return l*l
}

// function main
func main(){
    var value int
    fmt.Println("enter with value: ")
    fmt.Scan(&value)
    a := area(value)
    fmt.Println("area of a square: ", a)
}
