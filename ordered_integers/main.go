package main

import "fmt"
import "sort"

func orderedIntegers(){
	fmt.Println("Sorting integers...")
	sorting := []int{4, 10, 2, 3, 1}
	sort.Ints(sorting)
	fmt.Println(sorting)
}

func main(){
	orderedIntegers()
}
