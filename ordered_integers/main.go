package main

import "fmt"
import "sort"

// function ordered
func orderedIntegers(sorting[]int) {
	sort.Ints(sorting)
}

// function main
func main(){
	numbers := []int{4, 10, 2, 3, 1}
	orderedIntegers(numbers)
	fmt.Println("Sorting integers...")
	fmt.Println(numbers)
	bigger := numbers[len(numbers)-1]
	fmt.Println("Bigger number is: ", bigger)
}
