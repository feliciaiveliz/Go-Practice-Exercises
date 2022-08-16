package main

import (
	"fmt"
	"strconv"
	"strings"
)

// square all numbers from 0 to 'n'
// put all numbers into slice
// set count to 0
// iterate through the slice and count the number of 1s
// convert each number to string and get the count of 1s

func NbDig(n int, d int) (count int) {
	numbers := []int{}

	for i := 0; i <= n; i++ {
		numbers = append(numbers, i*i)
	}

	for _, n := range numbers {
		str := strconv.Itoa(n)
		count += strings.Count(str, strconv.Itoa(d))
	}

	fmt.Println(count)
	return
}

func main() {
	NbDig(25, 1) // 11
	NbDig(10, 1) // 4
}
