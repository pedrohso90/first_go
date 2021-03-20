package main

import (
	"fmt"
	"sort"
)

// function ordered and bigger number
func sortGetMaxArrayNumber(aSort []int) (int, []int) {
	sort.Ints(aSort)
	maxNumber := aSort[len(aSort)-1]
	return maxNumber, aSort
}

// function main
func main() {
	numbers := []int{4, 10, 2, 3, 1}

	var maxNumber int = 0

	maxNumber, numbers = sortGetMaxArrayNumber(numbers)

	fmt.Println("Integers sorted: ", numbers)

	fmt.Println("Max number: ", maxNumber)
}
