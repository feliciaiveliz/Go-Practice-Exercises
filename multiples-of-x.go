package main

import (
	"fmt"
)

// loop from 1 to 'n'
// multiply 'x' by current index number (starting at 1)
// add this number to result[index]
// return result

func CountBy(x, n int) []int {
	result := make([]int, n)

	for idx := 0; idx <= (n - 1); idx++ {
		result[idx] = x * (idx + 1)
	}

	fmt.Println(result)
	return result
}

func main() {
	CountBy(2, 5)
}
