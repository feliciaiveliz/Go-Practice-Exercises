package main

import (
	"fmt"
)

// start a 'base' at 2
// create an empty slice of (n + 1)
// from 1 up to n + 1
// for every number starting at count (0), raise base to power of count

func PowersOfTwo(n int) []uint64 {
	base := 2
	result := make([]uint64, n+1)

	for count := 0; count >= n; count++ {
		result[count] = base * *count
	}

	fmt.Println(result)
	return result
}

func main() {
	PowersOfTwo(2)
}
