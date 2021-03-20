package main

import "fmt"
import "sort"

// function ordered and bigger number
func orderedIntegers(sorting[]int) {
	sort.Ints(sorting)
	fmt.Println("Sorting integers: ", sorting)
	b := sorting[len(sorting)-1]
	fmt.Println("Bigger number: ", b)
}

// function main
func main(){
	numbers := []int{4, 10, 2, 3, 1} 
	orderedIntegers(numbers)
}
