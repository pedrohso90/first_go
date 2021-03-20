package main

import "fmt"
import "sort"

// function ordered
func orderedIntegers(sorting[]int) {
	sort.Ints(sorting)
}

func biggerNumber(bigger[]int) { 
	b := bigger[len(bigger)-1] 
	return b
}

// function main
func main(){
	numbers := []int{4, 10, 2, 3, 1} 
	var bigger int
	orderedIntegers(numbers)
	biggerNumber(numbers)
	fmt.Println("Sorting integers: ", numbers, "Bigger number: ", bigger)
}
