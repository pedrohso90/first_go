package main

import "fmt"
import "sort"

// function ordered
func orderedIntegers(sorting[]int) {
	sort.Ints(sorting)
}

func biggerNumber(bigger[]int) {
	bigger := numbers[len(numbers)-1] 
}

// function main
func main(){
	numbers := []int{4, 10, 2, 3, 1}
	orderedIntegers(numbers)
	biggerNumber(numbers)
	fmt.Println("Sorting integers: ", mumbers, "Bigger number: ", bigger)
}
