package main

import (
	"fmt"
)

// loop through numbers
// from first number, add one to currentCount
// if the next number is less than the current number, return false
// return true if all numbers have a difference of 1

func InAscOrder(numbers []int) bool {
	for idx := 0; idx < len(numbers) - 1; idx++ {
      if numbers[idx] > numbers[idx + 1] {
		  fmt.Println(false)
		  return false
	  }
	}

	fmt.Println(true)
    return true
}

func main() {
	InAscOrder([]int{1, 2, 4, 7, 19})         // true
	InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}) // false
}