package main

import (
	"fmt"
	"sort"
)

func sortArrays(array1, array2 []int) []int {
	result := make([]int, len(array1) + len(array2))

  for i := 0; i < len(array1); i++ {
    result[i] = array1[i]
  }

  j := len(array1)
  for i := 0; i < len(array2); i++ {
    result[j] = array2[i]
    j++
  }

  sort.Ints(result)
  fmt.Println(result)
  return result
}

func removeDuplicates(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, value := range slice {
    if _, 
	}

  return result
}

func MergeArrays(arr1, arr2 []int) []int {
  result := sortArrays(arr1, arr2)
	result = removeDuplicates(result)

	fmt.Println(result)
	return result
}

func main() {
	MergeArrays([]int{1,3,5,7,9}, []int{10,8,6,4,2})
}