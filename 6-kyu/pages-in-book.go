package main

import (
	"fmt"
	"strconv"
)

// input: number
// output: number of pages as digits
// rules:
// - # of pages calculated: digits summed from 0 to 'n'
// algorithm:
// - loop from 1 to 'n':
//   - for each number, convert the number to a string
//   - add that new string digit to 'digits'
// - '12345' if n = 5
// - find the number of pages it takes until we reach the summary number
// loop until the count is equal to the summary digit

func AmountOfPages(summary int) int {
	digits := ""
	var count int

	for count = 1; count < summary; count++ {
		digits += strconv.Itoa(count)
		if len(digits) == summary {
			break
		}
	}

	fmt.Println(count)
	return count
}

func main() {
	AmountOfPages(5)    // 5 -> "12345"
	AmountOfPages(25)   // 17
	AmountOfPages(7242) // 2087  -- getting 7243
	AmountOfPages(1095) // 401
	AmountOfPages(185)  // 97
	AmountOfPages(660)  // 256
}
