package main

import "fmt"
import "sort"

func orderedIntegers(sorting[]int) {
	sort.Ints(sorting)
}

func main(){
	numbers := []int{4, 10, 2, 3, 1}
	orderedIntegers(numbers)
	fmt.Println("Sorting integers...")
	fmt.Println(numbers)
}
