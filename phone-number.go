package main

import (
	"fmt"
	"strconv"
)

// loop 3 times, taking a slice each time
// add ( first
// then "123"
// add ") "
// then "456-"
// then "7890"

func CreatePhoneNumber(numbers [10]uint) string {
  result := "("

	for i := 0; i <= len(numbers) - 1; i++ {
    result += strconv.FormatUint(uint64(numbers[i]), 10)

		if i == 2 { break }
	}

	result += ") "

	for i := 3; i <= len(numbers) - 1; i++ {
		result += strconv.FormatUint(uint64(numbers[i]), 10)

		if i == 5 { break }
	}

	result += "-"

	for i := 6; i <= len(numbers) - 1; i++ {
    result += strconv.FormatUint(uint64(numbers[i]), 10)
  }

	fmt.Println(result)
	return result
}

func main() {
	CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // "(123) 456-7890"
}
