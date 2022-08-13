// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// set count to 0
// create a range from a to b inclusive
// 2 3 4 5 6 7
// iterate through slice and convert each int into binary
// iterate through slice once more and count the number of 1s in each binary
// return count

func RangeBitCount(a, b int) (result int) {
	numbers := []int{}

	for i := a; i <= b; i++ {
		numbers = append(numbers, i)
	}

	binary := make([]string, len(numbers))

	for idx, n := range numbers {
		binary[idx] = strconv.FormatInt(int64(n), 2)
	}

	for _, bin := range binary {
		result += strings.Count(bin, "1")
	}

	return
}

func main() {
	fmt.Println(RangeBitCount(2, 7)) // 11
}
