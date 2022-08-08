package main

import "fmt"

func SquareSum(numbers []int) (result int) {
	for _, value := range numbers {
		result += value * value
	}

	fmt.Println(result)
	return
}

func main() {
	SquareSum([]int{1, 2})
}
