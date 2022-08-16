package main

import (
	"sort"
)

// create a copy of the array (slice)
// sort the copied slice
// find the middle element's index
// return the number at index of original array

func Gimme(array [3]int) int {
	copy := []int{}

	for _, n := range array {
		copy = append(copy, n)
	}

	sort.Ints(copy)
	middle := copy[1]

	for idx, n := range array {
		if n == middle {
			return idx
		}
	}

	return 0
}

func main() {
	Gimme([3]int{2, 3, 1})   // 0
	Gimme([3]int{5, 10, 14}) // 1
}
