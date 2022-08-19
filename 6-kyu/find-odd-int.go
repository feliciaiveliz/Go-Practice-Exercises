package main 

import "fmt"

// create a map of numbers and the count of how many times it appears in slice
// iterate over map and return the number with a count of 1

func FindOdd(seq []int) int {
	counts := make(map[int]int)

	for _, n := range seq {
		counts[n] += 1
	}

	for number, count := range counts {
    if count == 1 { return number }
	}
}

func main() {
	FindOdd([]int{1,2,2,3,3,3,4,3,3,3,2,2,1}) // 4 -- 1 time
}
