package main

import (
	"sort"
)

// sort array -> a < b
// append to array arr[0] and arr[len(arr) - 1]

func MinMax(array []int) (result [2]int) {
	sort.Ints(array)

	result[0] = array[0]
	result[len(result)-1] = array[len(array)-1]

	return
}

func main() {
	MinMax([]int{1, 2, 3, 4, 5}) // [1, 5]
	MinMax([]int{5423, 6, 90, 13})
}
